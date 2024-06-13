package logic

import (
	"context"
	"dianshang/testapp/Pay/internal/svc"
	"dianshang/testapp/Pay/internal/types"
	"errors"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req *types.CreateOrderReq) (resp *types.CommonResply, err error) {
	// todo: add your logic here and delete this line
	// 从JWT令牌中提取用户信息
	username, ok := l.ctx.Value("userName").(string)
	if !ok {
		// 如果无法获取username，返回错误
		return nil, errors.New("username is missing from context")
	}
	return &types.CommonResply{
		Code:    200,
		Message: username + "成功创建购买" + strconv.Itoa(req.BuyQuantity) + "个ProductId：“" + strconv.Itoa(req.ProductId) + "”商品的订单",
		Data:    "",
	}, nil
}
