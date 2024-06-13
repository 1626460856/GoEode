package handler

import (
	"context"
	"dianshang/testapp/Pay/database"
	"dianshang/testapp/User/middleware"
	"encoding/json"
	"errors"
	"net/http"

	"dianshang/testapp/Pay/internal/logic"
	"dianshang/testapp/Pay/internal/svc"
	"dianshang/testapp/Pay/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type DeleteOrderMessage struct {
	OrderId int `json:"OrderId"`
}

// DeleteOrderHandler 删除订单
func DeleteOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteOrderReq
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
		order, _ := database.GetOrderById(database.ShopRedis2DB, req.OrderId)
		if order.UserName != username {
			httpx.ErrorCtx(r.Context(), w, errors.New("该用户不是订单所有者"))
			return
		}
		// 创建一个 DeleteOrderMessage 结构体实例
		msg := DeleteOrderMessage{
			OrderId: req.OrderId,
		}

		// 将 msg 转换为 JSON
		reqData, err := json.Marshal(msg)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 发送消息到 Kafka
		err = svcCtx.KafkaClient.SendMessage("DeleteOrder", reqData)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		// 将 username等信息 放入上下文中
		ctx := context.WithValue(r.Context(), "UserName", order.UserName)
		ctx = context.WithValue(ctx, "ProductName", order.ProductName)
		ctx = context.WithValue(ctx, "BuyQuantity", order.BuyQuantity)
		l := logic.NewDeleteOrderLogic(ctx, svcCtx)
		resp, err := l.DeleteOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
