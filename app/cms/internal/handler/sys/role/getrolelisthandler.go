package role

import (
	"net/http"

	"chenxi/app/cms/internal/logic/sys/role"
	"chenxi/app/cms/internal/svc"
	"chenxi/app/cms/internal/types"
	"chenxi/pkg/utils/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetRoleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRoleListRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, r, nil, err)
			return
		}

		l := role.NewGetRoleListLogic(r.Context(), svcCtx)
		resp, err := l.GetRoleList(&req)
		response.Response(w, r, resp, err)
	}
}
