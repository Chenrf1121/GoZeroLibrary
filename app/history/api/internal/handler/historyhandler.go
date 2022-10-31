package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"library/app/history/api/internal/logic"
	"library/app/history/api/internal/svc"
	"library/app/history/api/internal/types"
)

func historyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Hisreq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHistoryLogic(r.Context(), svcCtx)
		resp, err := l.History(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
