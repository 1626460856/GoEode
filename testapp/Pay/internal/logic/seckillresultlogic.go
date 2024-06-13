package logic

import (
	"context"
	"dianshang/testapp/Pay/database"
	"errors"
	"strconv"

	"dianshang/testapp/Pay/internal/svc"
	"dianshang/testapp/Pay/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeckillResultLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSeckillResultLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillResultLogic {
	return &SeckillResultLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SeckillResultLogic) SeckillResult(req *types.SeckillResultReq) (resp *types.SeckillResultResply, err error) {
	// 从JWT令牌中提取用户信息
	username, ok := l.ctx.Value("UserName").(string)
	if !ok {
		// 如果无法获取username，返回错误
		return nil, errors.New("username is missing from context")
	}
	// 从 Redis 获取所有键（商品 ID）
	ctx := context.Background()
	keys, err := database.ShopRedis2DB.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}

	// 初始化一个空切片来存储商品
	var orders []database.Order

	// 遍历键并获取每个商品的详细信息
	for _, key := range keys {
		id, err := strconv.Atoi(key)
		if err != nil {
			return nil, err
		}

		order, err := database.GetOrderById(database.ShopRedis2DB, id)
		if err != nil {
			return nil, err
		}
		// 检查订单的UserName是否等于UserName，订单的ProductID是否等于req。ProductId
		if order.UserName == username && order.ProductID == req.ProductId {
			// 将商品添加到切片中
			orders = append(orders, order)
		}
	}

	// 返回商品列表
	return &types.SeckillResultResply{
		Code:    200,
		Message: "成功获取商品",
		Data:    orders,
	}, nil
}
