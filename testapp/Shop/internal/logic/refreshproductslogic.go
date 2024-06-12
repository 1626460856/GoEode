package logic

import (
	"context"

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

func (l *RefreshProductsLogic) RefreshProducts(req *types.RefreshProductsReq) (resp *types.CommonResply, err error) {
	// todo: add your logic here and delete this line

	return
}
