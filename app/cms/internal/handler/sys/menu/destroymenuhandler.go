package menu

import (
	"net/http"

	"chenxi/app/cms/internal/logic/sys/menu"
	"chenxi/app/cms/internal/svc"
	"chenxi/app/cms/internal/types"
	"chenxi/pkg/utils/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DestroyMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DestroyMenuRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, r, nil, err)
			return
		}

		l := menu.NewDestroyMenuLogic(r.Context(), svcCtx)
		err := l.DestroyMenu(&req)
		response.Response(w, r, nil, err)
	}
}
