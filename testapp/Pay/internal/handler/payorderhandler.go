package handler

import (
	"net/http"

	"dianshang/testapp/Pay/internal/logic"
	"dianshang/testapp/Pay/internal/svc"
	"dianshang/testapp/Pay/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// PayOrderHandler 支付订单
func PayOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PayOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewPayOrderLogic(r.Context(), svcCtx)
		resp, err := l.PayOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
