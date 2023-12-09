package main

import (
	"chenxi/app/mqueue/rabbitmq/internal/config"
	"chenxi/app/mqueue/rabbitmq/internal/handler"
	"chenxi/app/mqueue/rabbitmq/internal/svc"
	"context"
	"flag"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/rabbitmq.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c)

	logx.DisableStat()
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	ctx := context.Background()
	svcContext := svc.NewServiceContext(c)
	serverGroup := handler.RegisterHandlers(ctx, svcContext)
	defer serverGroup.Stop()
	logx.Info("rabbitmq listener run ...")
	serverGroup.Start()
}
