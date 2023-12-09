package config

import (
	"chenxi/pkg/dao/mysql"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	CloudMysql mysql.Mysql
	Redis      redis.RedisConf
}
