package handler

import (
	"context"
	"dianshang/testapp/Pay/middleware"
	"errors"
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
		// 从上下文中获取 JWT 令牌

		username, ok := r.Context().Value(middleware.ContextKeyUsername).(string)
		if !ok {
			httpx.ErrorCtx(r.Context(), w, errors.New("username is missing from context"))
			return
		}
		// 将 username 放入上下文中
		ctx := context.WithValue(r.Context(), "UserName", username)

		l := logic.NewSeckillResultLogic(ctx, svcCtx)
		resp, err := l.SeckillResult(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
