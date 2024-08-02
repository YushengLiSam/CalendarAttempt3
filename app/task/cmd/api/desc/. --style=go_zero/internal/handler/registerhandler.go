package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"calendarAttempt3/app/task/cmd/api/desc/. --style=go_zero/internal/logic"
	"calendarAttempt3/app/task/cmd/api/desc/. --style=go_zero/internal/svc"
	"calendarAttempt3/app/task/cmd/api/desc/. --style=go_zero/internal/types"

)

func registerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
