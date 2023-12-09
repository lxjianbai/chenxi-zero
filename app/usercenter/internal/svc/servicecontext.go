package svc

import (
	"chenxi/app/usercenter/internal/config"
	"chenxi/pkg/dao/mysql"
	"chenxi/pkg/model/cloudmodel"
	"errors"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/syncx"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config     config.Config
	CloudDb    *gorm.DB
	Redis      *redis.Redis
	CloudModel *cloudmodel.CloudModel
	Cache      cache.Cache //缓存
}

func NewServiceContext(c config.Config) *ServiceContext {
	cloudDb, err := mysql.ConnectMysql(c.CloudMysql)
	if err != nil {
		panic("mysql init fail")
	}
	redis := redis.MustNewRedis(c.Redis.RedisConf)
	cache := cache.NewNode(redis, syncx.NewSingleFlight(), cache.NewStat("cache"), errors.New("err"))
	return &ServiceContext{
		Config:     c,
		CloudDb:    cloudDb,
		Redis:      redis,
		Cache:      cache,
		CloudModel: cloudmodel.NewCloudModel(cloudDb),
	}
}
