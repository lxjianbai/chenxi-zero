package login

import (
	"net/http"

	"chenxi/app/cms/internal/logic/sys/login"
	"chenxi/app/cms/internal/svc"
	"chenxi/app/cms/internal/types"
	"chenxi/pkg/utils/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, r, nil, err)
			return
		}

		l := login.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Response(w, r, resp, err)
	}
}
