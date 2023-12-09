package svc

import (
	"chenxi/app/cms/internal/config"
	"chenxi/pkg/dao/mysql"
	"chenxi/pkg/model/cloudmodel"
	"chenxi/pkg/model/gamemodel"
	"errors"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/syncx"
)

type ServiceContext struct {
	Config     config.Config
	Redis      *redis.Redis
	Cache      cache.Cache //缓存
	CloudModel *cloudmodel.CloudModel
	GameModel  *gamemodel.GameModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	cloudDb, err := mysql.ConnectMysql(c.CloudMysql)
	if err != nil {
		panic("mysql init fail")
	}

	gameDb, err := mysql.ConnectMysql(c.GameMysql)
	if err != nil {
		panic("mysql init fail")
	}

	redis := redis.MustNewRedis(c.Redis)
	cache := cache.NewNode(redis, syncx.NewSingleFlight(), cache.NewStat("cache"), errors.New("err"))

	cloudModel := cloudmodel.NewCloudModel(cloudDb)
	gameModel := gamemodel.NewGameModel(gameDb)

	return &ServiceContext{
		Config:     c,
		Cache:      cache,
		Redis:      redis,
		CloudModel: cloudModel,
		GameModel:  gameModel,
	}
}
