package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"calendarAttempt3/app/task/cmd/api/desc/. --style=go_zero/internal/logic"
	"calendarAttempt3/app/task/cmd/api/desc/. --style=go_zero/internal/svc"
	"calendarAttempt3/app/task/cmd/api/desc/. --style=go_zero/internal/types"

)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
