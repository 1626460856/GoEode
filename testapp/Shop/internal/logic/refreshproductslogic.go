package logic

import (
	"context"
	"dianshang/testapp/Shop/database"
	"strconv"

	"dianshang/testapp/Shop/internal/svc"
	"dianshang/testapp/Shop/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshProductsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshProductsLogic {
	return &RefreshProductsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshProductsLogic) RefreshProducts(req *types.RefreshProductsReq) (resp *types.RefreshProductsResply, err error) {
	// 从 Redis 获取所有键（商品 ID）
	ctx := context.Background()
	keys, err := database.ShopRedis1DB.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}

	// 初始化一个空切片来存储商品
	var products []types.Product

	// 遍历键并获取每个商品的详细信息
	for _, key := range keys {
		id, err := strconv.Atoi(key)
		if err != nil {
			return nil, err
		}

		product, err := database.GetProduct(database.ShopRedis1DB, id)
		if err != nil {
			return nil, err
		}

		// 将商品添加到切片中
		products = append(products, types.Product{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			Boss:        product.Boss,
		})
	}

	// 返回商品列表
	return &types.RefreshProductsResply{
		Code:    200,
		Message: "成功获取商品",
		Data:    products,
	}, nil
}
