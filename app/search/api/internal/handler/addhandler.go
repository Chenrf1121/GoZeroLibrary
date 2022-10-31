package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"library/app/search/api/internal/logic"
	"library/app/search/api/internal/svc"
	"library/app/search/api/internal/types"
)

func addHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddLogic(r.Context(), svcCtx)
		resp, err := l.Add(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
