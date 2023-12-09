package config

import (
	"chenxi/pkg/dao/mysql"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	CloudMysql mysql.Mysql
}
