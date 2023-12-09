package config

import (
	"chenxi/pkg/dao/mysql"
	"chenxi/pkg/queue/rabbitmq"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Redis            redis.RedisConf
	CloudMysql       mysql.Mysql
	TestListenerConf rabbitmq.RabbitListenerConf
}
