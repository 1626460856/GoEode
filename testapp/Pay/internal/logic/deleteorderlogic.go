package logic

import (
	"context"

	"dianshang/testapp/Pay/internal/svc"
	"dianshang/testapp/Pay/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderLogic {
	return &DeleteOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteOrderLogic) DeleteOrder(req *types.DeleteOrderReq) (resp *types.CommonResply, err error) {
	// todo: add your logic here and delete this line

	return
}
