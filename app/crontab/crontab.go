package main

import (
	"chenxi/app/crontab/internal/config"
	"chenxi/app/crontab/internal/job"
	"chenxi/app/crontab/internal/scheduler"
	"chenxi/app/crontab/internal/svc"
	"context"
	"flag"
	"os"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/crontab.yaml", "Specify the config file")

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

	go func() {
		scheduler.Register(ctx, svcContext)
		if err := svcContext.Scheduler.Run(); err != nil {
			logx.Errorf("!!!MqueueSchedulerErr!!!  run err:%+v", err)
			os.Exit(1)
		}
	}()

	mux := job.Register(ctx, svcContext)
	mux.Use(svc.LoggingMiddleware)
	if err := svcContext.AsynqServer.Run(mux); err != nil {
		logx.Errorf("!!!CronJobErr!!! run err:%+v", err)
		os.Exit(1)
	}
}
