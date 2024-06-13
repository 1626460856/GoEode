package logic

import (
	"context"
	"dianshang/testapp/Pay/database"
	"dianshang/testapp/Pay/internal/svc"
	"dianshang/testapp/Pay/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.GetOrderReq) (resp *types.Order, err error) {
	// todo: add your logic here and delete this line

	order, err := database.GetOrderById(database.ShopRedis2DB, req.OrderId)
	return &types.Order{
		OrderID:     order.OrderID,
		ProductID:   order.ProductID,
		ProductName: order.ProductName,
		Price:       order.Price,
		Boss:        order.Boss,
		BuyQuantity: order.BuyQuantity,
		UserName:    order.UserName,
		Coupon:      order.Coupon,
		OrderStatus: order.OrderStatus,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
	}, nil
}
