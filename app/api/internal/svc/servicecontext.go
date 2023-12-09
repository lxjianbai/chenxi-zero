package svc

import (
	"chenxi/app/api/internal/config"
	"chenxi/app/usercenter/usercenterclient"
	"chenxi/pkg/dao/mysql"
	"chenxi/pkg/model/cloudmodel"
	"chenxi/pkg/queue/nats"
	"errors"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/syncx"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config        config.Config
	CloudDb       *gorm.DB //云游db
	Redis         *redis.Redis
	Cache         cache.Cache            //缓存
	CloudModel    *cloudmodel.CloudModel // 云游model
	NatsProducer  *nats.Producer
	UserCenterRpc usercenterclient.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	cloudDb, err := mysql.ConnectMysql(c.CloudMysql)
	if err != nil {
		panic("mysql init fail")
	}

	redis := redis.MustNewRedis(c.Redis)
	cache := cache.NewNode(redis, syncx.NewSingleFlight(), cache.NewStat("cache"), errors.New("err"))

	cloudModel := cloudmodel.NewCloudModel(cloudDb)

	natsProducer, err := nats.NewProducer(&nats.NatsConfig{
		Host: c.Nats.Host,
	})
	if err != nil {
		panic("nats init fail")
	}

	return &ServiceContext{
		Config:        c,
		CloudDb:       cloudDb,
		Cache:         cache,
		CloudModel:    cloudModel,
		NatsProducer:  natsProducer,
		UserCenterRpc: usercenterclient.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
