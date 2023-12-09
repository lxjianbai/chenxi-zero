package role

import (
	"context"

	"chenxi/app/cms/internal/svc"
	"chenxi/app/cms/internal/types"
	"chenxi/pkg/utils/xerror"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleLogic) UpdateRole(req *types.SysRole) error {
	sysRole, err := l.svcCtx.GameModel.SysRoleModel.FindOne(l.ctx, uint64(req.Id))
	if err != nil {
		return xerror.NewDefaultError("角色不存在")
	}
	sysRole.MenuIds = req.MenuIds
	sysRole.Name = req.Name
	sysRole.Remark = req.Remark
	sysRole.UniqueKey = req.UniqueKey
	sysRole.Status = uint64(req.Status)
	err = l.svcCtx.GameModel.SysRoleModel.Update(l.ctx, nil, sysRole)
	if err != nil {
		return xerror.NewInternalError("角色更新失败")
	}
	return nil
}
