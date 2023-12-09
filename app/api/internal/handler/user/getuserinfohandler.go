package user

import (
	"net/http"

	"chenxi/app/api/internal/logic/user"
	"chenxi/app/api/internal/svc"
	"chenxi/app/api/internal/types"
	"chenxi/pkg/utils/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmptyRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, r, nil, err)
			return
		}

		l := user.NewGetUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfo(&req)
		response.Response(w, r, resp, err)
	}
}
