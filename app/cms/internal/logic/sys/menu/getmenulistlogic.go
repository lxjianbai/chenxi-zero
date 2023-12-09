package menu

import (
	"context"
	"sort"

	"chenxi/app/cms/internal/svc"
	"chenxi/app/cms/internal/types"
	"chenxi/pkg/model/gamemodel"
	"chenxi/pkg/utils/xerror"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListLogic {
	return &GetMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuListLogic) GetMenuList(req *types.GetMenuListRequest) (resp []*types.Menu, err error) {
	var menus []*gamemodel.SysMenu
	var model = l.svcCtx.GameModel.SysMenuModel.Builder(nil)
	if req.Title != "" {
		model.Where("title = ?", req.Title)
	}
	if req.Name != "" {
		model.Where("name = ?", req.Name)
	}
	err = model.Find(&menus).Error
	if err != nil {
		return nil, xerror.NewInternalError(err.Error())
	}
	sort.Slice(menus, func(i, j int) bool {
		return menus[i].Sort < menus[j].Sort
	})
	resp = svc.TreeMenu(menus, 0)
	return
}
