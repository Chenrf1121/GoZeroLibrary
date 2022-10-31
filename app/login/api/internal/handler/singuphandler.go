package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"library/app/login/api/internal/logic"
	"library/app/login/api/internal/svc"
	"library/app/login/api/internal/types"
)

func sing_upHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SignupReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSing_upLogic(r.Context(), svcCtx)
		resp, err := l.Sing_up(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
