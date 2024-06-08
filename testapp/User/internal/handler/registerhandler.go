package handler

import (
	"encoding/json"
	"net/http"

	"dianshang/testapp/User/internal/logic"
	"dianshang/testapp/User/internal/svc"
	"dianshang/testapp/User/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
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

		// 发送消息到 Kafka
		err = svcCtx.KafkaClient.SendMessage("testtopic", reqData)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 调用业务逻辑处理
		//创建一个新的 RegisterLogic 实例 l，并调用其 Register 方法来处理业务逻辑
		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
