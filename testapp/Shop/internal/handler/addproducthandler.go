package handler

import (
	"context"
	"dianshang/testapp/Shop/database"
	"encoding/json"
	"net/http"

	"dianshang/testapp/Shop/internal/logic"
	"dianshang/testapp/Shop/internal/svc"
	"dianshang/testapp/Shop/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddProductReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 将请求数据转换为 JSON
		reqData, err := json.Marshal(req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		// 从Redis2中获取用户信息
		ctx := context.Background()
		exists, err := database.UserRedis2DB.SIsMember(ctx, "UserName", req.Boss).Result()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		if !exists {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		// 发送消息到 Kafka
		err = svcCtx.KafkaClient.SendMessage("AddProduct", reqData)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewAddProductLogic(r.Context(), svcCtx)
		resp, err := l.AddProduct(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
