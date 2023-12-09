package main

import (
	"flag"

	"chenxi/app/api/internal/config"
	"chenxi/app/api/internal/handler"
	"chenxi/app/api/internal/middleware"
	"chenxi/app/api/internal/svc"
	"chenxi/pkg/utils/response"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.DisableStat()

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(response.JwtUnauthorizedResult))

	server.Use(middleware.NewCorsMiddleware().Handle)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	logx.Infof("Starting server at %s:%d...", c.Host, c.Port)
	server.Start()
}
