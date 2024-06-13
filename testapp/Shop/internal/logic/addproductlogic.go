package logic

import (
	"context"
	"strconv"

	"dianshang/testapp/Shop/internal/svc"
	"dianshang/testapp/Shop/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductLogic {
	return &AddProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddProductLogic) AddProduct(req *types.AddProductReq) (resp *types.CommonResply, err error) {
	// todo: add your logic here and delete this line

	return &types.CommonResply{
		Code:    200,
		Message: req.Boss + "成功添加" + strconv.Itoa(req.Stock) + "个“" + req.Name + "”商品到商店",
		Data:    "",
	}, nil
}
