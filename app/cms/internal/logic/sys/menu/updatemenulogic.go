package menu

import (
	"context"

	"chenxi/app/cms/internal/svc"
	"chenxi/app/cms/internal/types"
	"chenxi/pkg/utils/xerror"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMenuLogic) UpdateMenu(req *types.Menu) error {
	menu, err := l.svcCtx.GameModel.SysMenuModel.FindOne(l.ctx, uint64(req.Id))
	if err != nil {
		return xerror.NewInternalError(err.Error())
	}
	menu.Pid = uint64(req.Pid)
	menu.Path = req.Path
	menu.Type = uint64(req.Type)
	menu.Name = req.Name
	menu.Redirect = req.Redirect
	menu.Perms = req.Perms
	menu.Component = req.Component
	menu.Icon = req.Meta.Icon
	menu.Title = req.Meta.Title
	menu.Link = req.Meta.IsLink
	menu.IsAffix = uint64(req.Meta.IsAffix)
	menu.IsHide = uint64(req.Meta.IsHide)
	menu.IsFull = uint64(req.Meta.IsFull)
	menu.IsKeepalive = uint64(req.Meta.IsKeepAlive)
	menu.Sort = req.Sort

	err = l.svcCtx.GameModel.SysMenuModel.Update(l.ctx, nil, menu)
	if err != nil {
		return xerror.NewInternalError(err.Error())
	}
	return nil
}
