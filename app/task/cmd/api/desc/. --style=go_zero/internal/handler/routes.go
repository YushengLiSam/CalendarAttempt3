// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"calendarAttempt3/app/task/cmd/api/desc/. --style=go_zero/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
		{
			Method:  http.MethodPost,
			Path:    "/register",
			Handler: registerHandler(serverCtx),
		},
		{
			Method:  http.MethodPost,
			Path:    "/login",
			Handler: loginHandler(serverCtx),
		},
},  
rest.WithPrefix("/api/v1/user"),  
	)
}