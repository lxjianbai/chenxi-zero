package role

import (
	"net/http"

	"chenxi/app/cms/internal/logic/sys/role"
	"chenxi/app/cms/internal/svc"
	"chenxi/app/cms/internal/types"
	"chenxi/pkg/utils/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DestroyRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DestroyRoleRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, r, nil, err)
			return
		}

		l := role.NewDestroyRoleLogic(r.Context(), svcCtx)
		err := l.DestroyRole(&req)
		response.Response(w, r, nil, err)
	}
}
