package role

import (
	"context"

	"chenxi/app/cms/internal/svc"
	"chenxi/app/cms/internal/types"
	"chenxi/pkg/model/gamemodel"
	"chenxi/pkg/utils/xerror"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRoleLogic {
	return &AddRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddRoleLogic) AddRole(req *types.SysRole) error {
	var sysRole = new(gamemodel.SysRole)
	sysRole.Name = req.Name
	sysRole.UniqueKey = req.UniqueKey
	sysRole.Remark = req.Remark
	sysRole.Status = uint64(req.Status)
	sysRole.MenuIds = req.MenuIds

	err := l.svcCtx.GameModel.SysRoleModel.Insert(l.ctx, nil, sysRole)
	if err != nil {
		return xerror.NewInternalError("角色新增失败")
	}
	return nil
}
