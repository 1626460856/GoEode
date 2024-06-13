package handler

import (
	"dianshang/testapp/Pay/database"
	"dianshang/testapp/User/middleware"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"dianshang/testapp/Pay/internal/logic"
	"dianshang/testapp/Pay/internal/svc"
	"dianshang/testapp/Pay/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// GetOrderHandler 获取订单
func GetOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// 从 URL 查询参数中获取 id，不然直接?key=value会无法查找
		idStr := r.URL.Query().Get("orderId")
		if idStr == "" {
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("id is missing"))
			return
		}

		OrderId, err := strconv.Atoi(idStr)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("invalid id"))
			return
		}

		req := types.GetOrderReq{ // 传入请求参数
			OrderId: OrderId,
		}
		// 从上下文中获取 JWT 令牌

		username, ok := r.Context().Value(middleware.ContextKeyUsername).(string)
		if !ok {
			httpx.ErrorCtx(r.Context(), w, errors.New("username is missing from context1"))
			return
		}
		order, _ := database.GetOrderById(database.ShopRedis2DB, req.OrderId)
		if order.UserName != username {
			httpx.ErrorCtx(r.Context(), w, errors.New("该用户不是订单所有者"))
			return
		}

		l := logic.NewGetOrderLogic(r.Context(), svcCtx)
		resp, err := l.GetOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
