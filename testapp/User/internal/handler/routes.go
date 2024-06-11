// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"dianshang/testapp/User/middleware"
	"net/http"

	"dianshang/testapp/User/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/account"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/account/getuserinfo",
				Handler: middleware.JWTAuthMiddleware(getuserInfoHandler(serverCtx)),
			},
		},
	)
}
