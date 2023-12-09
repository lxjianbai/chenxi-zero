package main

import (
	"flag"
	"fmt"
	"time"

	"chenxi/app/cms/internal/config"
	"chenxi/app/cms/internal/handler"
	"chenxi/app/cms/internal/middleware"
	"chenxi/app/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/cms.yaml", "the config file")

func main() {
	flag.Parse()

	time.LoadLocation("Asia/Shanghai")
	var c config.Config
	conf.MustLoad(*configFile, &c)

	logx.DisableStat()
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	server.Use(middleware.NewCorsMiddleware().Handle)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
