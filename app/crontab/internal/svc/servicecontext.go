package svc

import (
	"chenxi/app/crontab/internal/config"
	"chenxi/pkg/dao/mysql"
	"chenxi/pkg/model/cloudmodel"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/syncx"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config      config.Config
	Scheduler   *asynq.Scheduler
	AsynqServer *asynq.Server
	CloudDb     *gorm.DB
	Redis       *redis.Redis
	Cache       cache.Cache //缓存
	CloudModel  *cloudmodel.CloudModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	cloudDb, err := mysql.ConnectMysql(c.CloudMysql)
	if err != nil {
		panic("mysql init fail")
	}

	redis := redis.MustNewRedis(c.Redis)
	cache := cache.NewNode(redis, syncx.NewSingleFlight(), cache.NewStat("cache"), errors.New("err"))

	cloudModel := cloudmodel.NewCloudModel(cloudDb)
	return &ServiceContext{
		Config:      c,
		Scheduler:   newScheduler(c),
		AsynqServer: newAsynqServer(c),
		CloudDb:     cloudDb,
		Cache:       cache,
		CloudModel:  cloudModel,
	}
}

func newScheduler(c config.Config) *asynq.Scheduler {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return asynq.NewScheduler(
		asynq.RedisClientOpt{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass,
		}, &asynq.SchedulerOpts{
			Location: location,
			EnqueueErrorHandler: func(task *asynq.Task, opts []asynq.Option, err error) {
				fmt.Printf("Scheduler EnqueueErrorHandler <<<<<<<===>>>>> err : %+v , task : %+v", err, task)
			},
		})
}

func newAsynqServer(c config.Config) *asynq.Server {
	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass},
		asynq.Config{
			IsFailure: func(err error) bool {
				fmt.Printf("asynq server exec task IsFailure ======== >>>>>>>>>>>  err : %+v \n", err)
				return true
			},
			Concurrency: 20, //max concurrent process job task num
		},
	)
}

// loggingMiddleware 记录任务日志中间件
func LoggingMiddleware(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		start := time.Now()
		logx.Infof("Start processing %q", t.Type())
		err := h.ProcessTask(ctx, t)
		if err != nil {
			return err
		}
		logx.Infof("Finished processing %q: Elapsed Time = %v", t.Type(), time.Since(start))
		return nil
	})
}
