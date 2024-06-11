package handler

import (
	"context"
	"dianshang/testapp/User/middleware"
	"errors"
	"net/http"

	"dianshang/testapp/User/internal/logic"
	"dianshang/testapp/User/internal/svc"
	"dianshang/testapp/User/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getuserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
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
		ctx := context.WithValue(r.Context(), "username", username)
		l := logic.NewGetuserInfoLogic(ctx, svcCtx)
		resp, err := l.GetuserInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
