package response

import (
	"chenxi/pkg/utils/xerror"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, r *http.Request, resp interface{}, err error) {
	var body Body
	if err != nil {
		switch e := err.(type) {
		case *xerror.CodeError: //业务输出错误
			body.Code = e.Code
			body.Message = e.Message
			body.Data = e.Data
		default: //内部错误
			body.Code = xerror.InternalCode
			body.Message = err.Error()
		}
	} else {
		body.Code = xerror.Success
		body.Message = "success"
		body.Data = resp
	}
	httpx.OkJsonCtx(r.Context(), w, body)
}

func JwtUnauthorizedResult(w http.ResponseWriter, r *http.Request, err error) {
	Response(w, r, nil, xerror.NewCodeError(xerror.AuthCode, "鉴权失败"))
}
