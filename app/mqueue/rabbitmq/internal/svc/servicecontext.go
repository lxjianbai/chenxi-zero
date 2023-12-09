package svc

import (
	"chenxi/app/mqueue/rabbitmq/internal/config"
	"chenxi/pkg/dao/mysql"
	"chenxi/pkg/model/cloudmodel"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config     config.Config
	CloudDb    *gorm.DB
	Redis      *redis.Redis
	CloudModel *cloudmodel.CloudModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	cloudDb, err := mysql.ConnectMysql(c.CloudMysql)
	if err != nil {
		panic("mysql init fail")
	}
	svc := &ServiceContext{
		Config:     c,
		CloudDb:    cloudDb,
		Redis:      redis.MustNewRedis(c.Redis),
		CloudModel: cloudmodel.NewCloudModel(cloudDb),
	}

	return svc
}
