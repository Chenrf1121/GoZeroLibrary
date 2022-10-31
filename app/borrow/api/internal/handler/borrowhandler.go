package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"library/app/borrow/api/internal/logic"
	"library/app/borrow/api/internal/svc"
	"library/app/borrow/api/internal/types"
)

func borrowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BorrowReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewBorrowLogic(r.Context(), svcCtx)
		resp, err := l.Borrow(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
