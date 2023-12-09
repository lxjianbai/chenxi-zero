package main

import (
	"chenxi/app/mqueue/nats/internal/config"
	"chenxi/app/mqueue/nats/internal/logic"
	"chenxi/app/mqueue/nats/internal/svc"
	"context"
	"flag"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/nats.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c)

	logx.DisableStat()
	// log、prometheus、trace、metricsUrl.
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()

	nats := logic.NewNatsSubscribe(ctx, svcContext)
	nats.Start()
}
