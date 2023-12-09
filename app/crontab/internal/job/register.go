package job

import (
	"chenxi/app/crontab/internal/entity"
	"chenxi/app/crontab/internal/svc"
	"context"

	"github.com/hibiken/asynq"
)

// register job
func Register(ctx context.Context, svcCtx *svc.ServiceContext) *asynq.ServeMux {
	mux := asynq.NewServeMux()
	//scheduler job
	mux.Handle(entity.TestEntity, NewTestCrontab(ctx, svcCtx))

	return mux
}
