package logic

import (
	"context"

	"dianshang/testapp/Pay/internal/svc"
	"dianshang/testapp/Pay/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UseCouponLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUseCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UseCouponLogic {
	return &UseCouponLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UseCouponLogic) UseCoupon(req *types.UseCouponReq) (resp *types.CommonResply, err error) {
	// todo: add your logic here and delete this line

	return
}
