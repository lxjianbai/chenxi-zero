package role

import (
	"net/http"

	"chenxi/app/cms/internal/logic/sys/role"
	"chenxi/app/cms/internal/svc"
	"chenxi/app/cms/internal/types"
	"chenxi/pkg/utils/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysRole
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, r, nil, err)
			return
		}

		l := role.NewUpdateRoleLogic(r.Context(), svcCtx)
		err := l.UpdateRole(&req)
		response.Response(w, r, nil, err)
	}
}
