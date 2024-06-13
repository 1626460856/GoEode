package handler

import (
	"net/http"

	"dianshang/testapp/Pay/internal/logic"
	"dianshang/testapp/Pay/internal/svc"
	"dianshang/testapp/Pay/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// SeckillResultHandler 查看秒杀结果
func SeckillResultHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SeckillResultReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSeckillResultLogic(r.Context(), svcCtx)
		resp, err := l.SeckillResult(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
