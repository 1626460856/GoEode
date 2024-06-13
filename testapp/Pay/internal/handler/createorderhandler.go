package handler

import (
	"net/http"

	"dianshang/testapp/Pay/internal/logic"
	"dianshang/testapp/Pay/internal/svc"
	"dianshang/testapp/Pay/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// CreateOrderHandler 创建订单
func CreateOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateOrderLogic(r.Context(), svcCtx)
		resp, err := l.CreateOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
