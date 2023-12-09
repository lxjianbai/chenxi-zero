package role

import (
	"context"

	"chenxi/app/cms/internal/svc"
	"chenxi/app/cms/internal/types"
	"chenxi/pkg/utils/xerror"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeRoleStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeRoleStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeRoleStatusLogic {
	return &ChangeRoleStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeRoleStatusLogic) ChangeRoleStatus(req *types.ChangeStatusRequest) error {
	sysRole, err := l.svcCtx.GameModel.SysRoleModel.FindOne(l.ctx, uint64(req.Id))
	if err != nil {
		return xerror.NewDefaultError("角色不存在")
	}
	sysRole.Status = uint64(req.Status)

	err = l.svcCtx.GameModel.SysRoleModel.Update(l.ctx, nil, sysRole)
	if err != nil {
		return xerror.NewInternalError("角色状态更改失败")
	}
	return nil
}
