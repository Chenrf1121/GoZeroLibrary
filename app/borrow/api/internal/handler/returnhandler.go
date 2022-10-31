package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"library/app/borrow/api/internal/logic"
	"library/app/borrow/api/internal/svc"
	"library/app/borrow/api/internal/types"
)

func returnHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReturnReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewReturnLogic(r.Context(), svcCtx)
		resp, err := l.Return(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
