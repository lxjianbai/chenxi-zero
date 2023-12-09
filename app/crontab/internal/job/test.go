package job

import (
	"chenxi/app/crontab/internal/svc"
	"context"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

type TestCrontab struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTestCrontab(ctx context.Context, svcCtx *svc.ServiceContext) *TestCrontab {
	return &TestCrontab{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (c *TestCrontab) ProcessTask(context.Context, *asynq.Task) error {
	return nil
}
