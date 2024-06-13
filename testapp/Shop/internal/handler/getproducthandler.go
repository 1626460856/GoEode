package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"dianshang/testapp/Shop/internal/logic"
	"dianshang/testapp/Shop/internal/svc"
	"dianshang/testapp/Shop/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从 URL 查询参数中获取 id，不然直接?key=value会无法查找
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("id is missing"))
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("invalid id"))
			return
		}

		req := types.GetProductReq{ // 传入请求参数
			Id: id,
		}

		l := logic.NewGetProductLogic(r.Context(), svcCtx)
		resp, err := l.GetProduct(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
