package config

import (
	"chenxi/pkg/dao/mysql"
	"chenxi/pkg/types"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Redis      redis.RedisConf
	CloudMysql mysql.Mysql
	Nats       types.NatsConf
	Auth       struct {
		AccessSecret string
		AccessExpire int64
	}
	UsercenterRpcConf zrpc.RpcClientConf
}
