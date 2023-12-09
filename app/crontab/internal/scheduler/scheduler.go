package scheduler

import (
	"chenxi/app/crontab/internal/entity"
	"chenxi/app/crontab/internal/svc"
	"context"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

func Register(ctx context.Context, svc *svc.ServiceContext) {
	testScheduler(ctx, svc)
}

func testScheduler(ctx context.Context, svc *svc.ServiceContext) {
	task := asynq.NewTask(entity.TestEntity, nil)
	// every one minute exec
	entryID, err := svc.Scheduler.Register("*/1 * * * *", task)
	if err != nil {
		logx.WithContext(ctx).Errorf("!!!CrontabSchedulerErr!!! ====>【testScheduler】registered err:%+v , task:%+v", err, task)
	}
	logx.Infof("【testScheduler】registered an entry: %q", entryID)
}
