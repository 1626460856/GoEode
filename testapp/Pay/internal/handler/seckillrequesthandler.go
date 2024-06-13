package handler

import (
	"net/http"

	"dianshang/testapp/Pay/internal/logic"
	"dianshang/testapp/Pay/internal/svc"
	"dianshang/testapp/Pay/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// SeckillRequestHandler 申请秒杀请求
func SeckillRequestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SeckillRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSeckillRequestLogic(r.Context(), svcCtx)
		resp, err := l.SeckillRequest(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
