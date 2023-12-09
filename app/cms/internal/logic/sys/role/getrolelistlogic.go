package role

import (
	"context"

	"chenxi/app/cms/internal/svc"
	"chenxi/app/cms/internal/types"
	"chenxi/pkg/model"
	"chenxi/pkg/model/gamemodel"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleListLogic {
	return &GetRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleListLogic) GetRoleList(req *types.GetRoleListRequest) (resp *types.GetRoleListResponse, err error) {
	var query = l.svcCtx.GameModel.SysRoleModel.Builder(nil)
	if req.Name != "" {
		query.Where("name = ?", req.Name)
	}
	if req.UniqueKey != "" {
		query.Where("unique_key = ?", req.UniqueKey)
	}
	var roles []*gamemodel.SysRole
	page := model.Paging(&model.Param{
		DB:       query,
		PageNum:  int(req.PageNum),
		PageSize: int(req.PageSize),
	}, &roles)

	var list []*types.SysRole
	resp = new(types.GetRoleListResponse)
	resp.PageNum = req.PageNum
	resp.PageSize = req.PageSize
	resp.Total = page.Total
	resp.List = model.PaginatorToApi[[]*types.SysRole](page, list)

	return
}
