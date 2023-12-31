// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	syslogin "chenxi/app/cms/internal/handler/sys/login"
	sysmenu "chenxi/app/cms/internal/handler/sys/menu"
	sysrole "chenxi/app/cms/internal/handler/sys/role"
	sysuser "chenxi/app/cms/internal/handler/sys/user"
	"chenxi/app/cms/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: syslogin.LoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/sys/menu/list",
				Handler: sysmenu.GetMenuListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/sys/menu/add",
				Handler: sysmenu.AddMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/sys/menu/update",
				Handler: sysmenu.UpdateMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/sys/menu/destroy",
				Handler: sysmenu.DestroyMenuHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/menus",
				Handler: sysuser.GetUserMenusHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/sys/role/list",
				Handler: sysrole.GetRoleListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/sys/role/add",
				Handler: sysrole.AddRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/sys/role/update",
				Handler: sysrole.UpdateRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/sys/role/destroy",
				Handler: sysrole.DestroyRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/sys/role/status",
				Handler: sysrole.ChangeRoleStatusHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
