package logic

import (
	"context"
	"errors"
	"strconv"

	"dianshang/testapp/Pay/internal/svc"
	"dianshang/testapp/Pay/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeckillRequestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSeckillRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillRequestLogic {
	return &SeckillRequestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SeckillRequestLogic) SeckillRequest(req *types.SeckillRequest) (resp *types.CommonResply, err error) {
	// todo: add your logic here and delete this line
	// 从JWT令牌中提取用户信息
	username, ok := l.ctx.Value("UserName").(string)
	if !ok {
		// 如果无法获取username，返回错误
		return nil, errors.New("username is missing from context")
	}
	return &types.CommonResply{
		Code:    200,
		Message: username + "成功将 “秒杀" + strconv.Itoa(req.BuyQuantity) + "个“ProductId:" + strconv.Itoa(req.ProductId) + "”商品” 的请求发送，请稍后查看秒杀结果",
		Data:    "",
	}, nil
}
