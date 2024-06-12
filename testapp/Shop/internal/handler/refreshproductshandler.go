package handler

import (
	"net/http"

	"dianshang/testapp/Shop/internal/logic"
	"dianshang/testapp/Shop/internal/svc"
	"dianshang/testapp/Shop/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RefreshProductsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RefreshProductsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRefreshProductsLogic(r.Context(), svcCtx)
		resp, err := l.RefreshProducts(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
