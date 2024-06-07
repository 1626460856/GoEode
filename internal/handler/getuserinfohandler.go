package handler

import (
	"net/http"

	"dianshang/internal/logic"
	"dianshang/internal/svc"
	"dianshang/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getuserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetuserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetuserInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
