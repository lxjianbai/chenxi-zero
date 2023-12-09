package user

import (
	"net/http"

	"chenxi/app/cms/internal/logic/sys/user"
	"chenxi/app/cms/internal/svc"
	"chenxi/pkg/utils/response"
)

func GetUserMenusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserMenusLogic(r.Context(), svcCtx)
		resp, err := l.GetUserMenus()
		response.Response(w, r, resp, err)
	}
}
