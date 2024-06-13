// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"dianshang/testapp/Pay/middleware"
	"net/http"

	"dianshang/testapp/Pay/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/pay/order/create",
				Handler: middleware.JWTAuthMiddleware(CreateOrderHandler(serverCtx)),
			},
			{
				Method:  http.MethodGet,
				Path:    "/pay/order/get",
				Handler: middleware.JWTAuthMiddleware(GetOrderHandler(serverCtx)),
			},
			{
				Method:  http.MethodPost,
				Path:    "/pay/order/payment",
				Handler: middleware.JWTAuthMiddleware(PayOrderHandler(serverCtx)),
			},
			{
				Method:  http.MethodPost,
				Path:    "/pay/order/usecoupon",
				Handler: middleware.JWTAuthMiddleware(UseCouponHandler(serverCtx)),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/pay/order/delete",
				Handler: middleware.JWTAuthMiddleware(DeleteOrderHandler(serverCtx)),
			},
			{
				Method:  http.MethodPost,
				Path:    "/pay/seckill/request",
				Handler: middleware.JWTAuthMiddleware(SeckillRequestHandler(serverCtx)),
			},
			{
				Method:  http.MethodPost,
				Path:    "/pay/seckill/result",
				Handler: middleware.JWTAuthMiddleware(SeckillResultHandler(serverCtx)),
			},
		},
	)
}