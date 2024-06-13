package logic

import (
	"context"
	"errors"
	"strconv"

	"dianshang/testapp/Pay/internal/svc"
	"dianshang/testapp/Pay/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPayOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayOrderLogic {
	return &PayOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PayOrderLogic) PayOrder(req *types.PayOrderReq) (resp *types.CommonResply, err error) {
	// todo: add your logic here and delete this line
	// 从JWT令牌中提取用户信息
	username, ok := l.ctx.Value("UserName").(string)
	productName, ok := l.ctx.Value("ProductName").(string)
	buyQuantity, ok := l.ctx.Value("BuyQuantity").(int)

	if !ok {
		// 如果无法获取username，返回错误
		return nil, errors.New("username is missing from context")
	}
	return &types.CommonResply{
		Code:    200,
		Message: username + "成功支付购买" + strconv.Itoa(buyQuantity) + "个“" + productName + "”商品的订单",
		Data:    "",
	}, nil
}
