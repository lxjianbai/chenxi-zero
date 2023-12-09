package user

import (
	"context"
	"encoding/json"
	"strings"

	"chenxi/app/cms/internal/svc"
	"chenxi/app/cms/internal/types"
	"chenxi/pkg/cpb"
	"chenxi/pkg/global/adminconst"
	"chenxi/pkg/model/gamemodel"
	"chenxi/pkg/utils/xerror"

	"github.com/gookit/goutil"
	"github.com/gookit/goutil/arrutil"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserMenusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserMenusLogic {
	return &GetUserMenusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserMenusLogic) GetUserMenus() (resp []*types.Menu, err error) {
	var uid = l.svcCtx.GetUidFromCtx(l.ctx)
	// 获取用户角色
	sysUser, err := l.svcCtx.GameModel.SysUserModel.FindOne(l.ctx, uid)
	if err != nil {
		return nil, xerror.NewDefaultError("用户不存在")
	}
	roles, err := l.getUserRoles(sysUser)
	if err != nil {
		return nil, xerror.NewInternalError("角色错误")
	}
	menus, err := l.getMenus(roles)
	if err != nil {
		return nil, xerror.NewInternalError("菜单获取失败")
	}

	resp = svc.TreeMenu(menus, 0)
	return
}

func (l *GetUserMenusLogic) getUserRoles(sysUser *gamemodel.SysUser) ([]*gamemodel.SysRole, error) {
	if sysUser.RoleIds == "" {
		return nil, xerror.NewDefaultError("角色错误,请联系管理员")
	}
	var roles []int64
	err := json.Unmarshal([]byte(sysUser.RoleIds), &roles)
	if err != nil {
		return nil, xerror.NewInternalError(err.Error())
	}
	var sysRole = make([]*gamemodel.SysRole, 0)

	if err = l.svcCtx.GameModel.SysRoleModel.Builder(nil).Where("id in ? and status = ?", roles, cpb.StatusSwitchEnum_StatusSwitch_Open).Find(&sysRole).Error; err != nil {
		return nil, err
	}
	return sysRole, nil
}

func (l *GetUserMenusLogic) getMenus(sysRoles []*gamemodel.SysRole) ([]*gamemodel.SysMenu, error) {
	var menus = make([]*gamemodel.SysMenu, 0)
	for _, role := range sysRoles {
		if role.Id == adminconst.SysSuperRoleId {
			if err := l.svcCtx.GameModel.SysMenuModel.Builder(nil).
				Where("type <> ? ", cpb.SysMenuTypeEnum_SysMenutype_Button).Find(&menus).Error; err != nil {
				return nil, err
			}
			return menus, nil
		}
	}
	var allPerms []int64
	for _, role := range sysRoles {
		var perms []int64
		for _, v := range strings.Split(role.MenuIds, ",") {
			permId, _ := goutil.ToInt64(v)
			perms = append(perms, permId)
		}
		allPerms = append(allPerms, perms...)
	}
	// 去重
	var menuIds = arrutil.Unique[int64](allPerms)
	err := l.svcCtx.GameModel.SysMenuModel.
		Builder(nil).
		Where("id in ?", menuIds, cpb.StatusSwitchEnum_StatusSwitch_Open).
		Find(&menus).
		Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}
