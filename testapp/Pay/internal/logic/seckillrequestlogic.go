package logic

import (
	"context"

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

	return
}
