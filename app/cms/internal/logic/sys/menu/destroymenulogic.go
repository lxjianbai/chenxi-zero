package menu

import (
	"context"

	"chenxi/app/cms/internal/svc"
	"chenxi/app/cms/internal/types"
	"chenxi/pkg/utils/xerror"

	"github.com/zeromicro/go-zero/core/logx"
)

type DestroyMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDestroyMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DestroyMenuLogic {
	return &DestroyMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DestroyMenuLogic) DestroyMenu(req *types.DestroyMenuRequest) error {
	err := l.svcCtx.GameModel.SysMenuModel.Delete(l.ctx, nil, uint64(req.Id))
	if err != nil {
		return xerror.NewInternalError(err.Error())
	}
	return nil
}
