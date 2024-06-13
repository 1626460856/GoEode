package logic

import (
	"context"
	"dianshang/testapp/Shop/database"

	"dianshang/testapp/Shop/internal/svc"
	"dianshang/testapp/Shop/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductLogic {
	return &GetProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductLogic) GetProduct(req *types.GetProductReq) (resp *types.Product, err error) {
	// todo: add your logic here and delete this line

	product, err := database.GetProduct(database.ShopRedis1DB, req.Id)
	if err != nil {
		return nil, err
	}

	return &types.Product{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Boss:        product.Boss,
	}, nil
}
