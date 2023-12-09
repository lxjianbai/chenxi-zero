package role

import (
	"net/http"

	"chenxi/app/cms/internal/logic/sys/role"
	"chenxi/app/cms/internal/svc"
	"chenxi/app/cms/internal/types"
	"chenxi/pkg/utils/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysRole
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, r, nil, err)
			return
		}

		l := role.NewAddRoleLogic(r.Context(), svcCtx)
		err := l.AddRole(&req)
		response.Response(w, r, nil, err)
	}
}
