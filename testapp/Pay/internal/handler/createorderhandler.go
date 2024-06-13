package handler

import (
	"context"
	"dianshang/testapp/User/middleware"
	"encoding/json"
	"errors"
	"net/http"

	"dianshang/testapp/Pay/internal/logic"
	"dianshang/testapp/Pay/internal/svc"
	"dianshang/testapp/Pay/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type CreateOrderMessage struct { //发送kafka消息
	// 订单id  mysql自增键
	ProductID int `json:"productId"` // 商品id传入获得
	// 商品名称kafka之后查询
	//商品价格kafka之后查询
	// 商家kafka之后查询
	BuyQuantity int    `json:"buyQuantity"` // 购买商品数量
	UserName    string `json:"userName"`    // 购买者这个通过传入的token解析获得
	// 优惠券,创建时默认为1
	// 订单状态 有三种状态，“unpaid”为未支付，“paying”为支付中，“paid”为已支付，创建的时候默认未支付
	// 创建时间kafka之后
	// 更新时间kafka之后
}

// CreateOrderHandler 创建订单
func CreateOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrderReq
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
		// 创建一个 CreateOrderMessage 结构体实例
		msg := CreateOrderMessage{
			ProductID:   req.ProductId,
			BuyQuantity: req.BuyQuantity,
			UserName:    username,
		}

		// 将 msg 转换为 JSON
		reqData, err := json.Marshal(msg)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 发送消息到 Kafka
		err = svcCtx.KafkaClient.SendMessage("CreateOrder", reqData)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		// 将 username 放入上下文中
		ctx := context.WithValue(r.Context(), "userName", username)

		l := logic.NewCreateOrderLogic(ctx, svcCtx)
		resp, err := l.CreateOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
