package role

import (
	"context"

	"chenxi/app/cms/internal/svc"
	"chenxi/app/cms/internal/types"
	"chenxi/pkg/utils/xerror"

	"github.com/zeromicro/go-zero/core/logx"
)

type DestroyRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDestroyRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DestroyRoleLogic {
	return &DestroyRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DestroyRoleLogic) DestroyRole(req *types.DestroyRoleRequest) error {
	err := l.svcCtx.GameModel.SysRoleModel.Delete(l.ctx, nil, uint64(req.Id))
	if err != nil {
		return xerror.NewInternalError("角色删除失败")
	}
	return nil
}
