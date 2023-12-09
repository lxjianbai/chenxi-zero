package logic

import (
	"chenxi/app/mqueue/nats/internal/svc"
	"chenxi/pkg/global/constants"
	"chenxi/pkg/queue/nats"
	"context"
)

type NatsSubscribe struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNatsSubscribe(ctx context.Context, svcCtx *svc.ServiceContext) *NatsSubscribe {
	return &NatsSubscribe{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NatsSubscribe) Start() {
	natsConf := &nats.NatsConfig{
		Host: l.svcCtx.Config.Nats.Host,
	}
	nats := nats.MustNewConsumerManager(natsConf, []*nats.ConsumerQueue{
		{
			QueueName: constants.Queue.ActiveLevel.QueueName,
			Subject:   constants.Queue.ActiveLevel.Subject,
			Consumer:  NewActiveLevel(l.ctx, l.svcCtx),
		},
	})

	nats.Start()
}
