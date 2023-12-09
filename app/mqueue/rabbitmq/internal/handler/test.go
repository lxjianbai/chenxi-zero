package handler

import (
	"chenxi/app/mqueue/rabbitmq/internal/svc"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestRabbitMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTestRabbitMq(ctx context.Context, svcCtx *svc.ServiceContext) *TestRabbitMq {
	return &TestRabbitMq{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (r *TestRabbitMq) Consume(message string) error {
	r.Logger.Info(message)
	return nil
}
