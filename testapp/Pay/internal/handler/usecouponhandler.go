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

type UseCouponMessage struct {
	Coupon  float64 `json:"Coupon"`
	OrderId int     `json:"OrderId"`
}

// UseCouponHandler 使用优惠券并且改变支付状态为正在支付
func UseCouponHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UseCouponReq
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
		// 创建一个 UserCouponMessage 结构体实例
		msg := UseCouponMessage{
			OrderId: req.OrderId,
			Coupon:  req.Coupon,
		}

		// 将 msg 转换为 JSON
		reqData, err := json.Marshal(msg)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 发送消息到 Kafka
		err = svcCtx.KafkaClient.SendMessage("UseCoupon", reqData)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		// 将 username等信息 放入上下文中
		ctx := context.WithValue(r.Context(), "UserName", order.UserName)
		ctx = context.WithValue(ctx, "ProductName", order.ProductName)
		ctx = context.WithValue(ctx, "BuyQuantity", order.BuyQuantity)
		l := logic.NewUseCouponLogic(ctx, svcCtx)
		resp, err := l.UseCoupon(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
