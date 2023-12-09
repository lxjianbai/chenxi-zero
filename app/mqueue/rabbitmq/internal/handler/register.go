package handler

import (
	"chenxi/app/mqueue/rabbitmq/internal/svc"
	"chenxi/pkg/queue/rabbitmq"
	"context"

	"github.com/zeromicro/go-zero/core/queue"
	"github.com/zeromicro/go-zero/core/service"
)

func RegisterHandlers(ctx context.Context, svcCtx *svc.ServiceContext) *service.ServiceGroup {
	serverGroup := service.NewServiceGroup()
	for _, listener := range makeListener(ctx, svcCtx) {
		serverGroup.Add(listener)
	}
	return serverGroup
}

func makeListener(ctx context.Context, svcCtx *svc.ServiceContext) []queue.MessageQueue {
	var listeners = make([]queue.MessageQueue, 0)
	testListener := rabbitmq.MustNewListener(svcCtx.Config.TestListenerConf, NewTestRabbitMq(ctx, svcCtx))
	listeners = append(listeners, testListener)
	return listeners
}
