package logic

import (
	"context"

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

func (l *SeckillResultLogic) SeckillResult(req *types.SeckillResultReq) (resp *types.CommonResply, err error) {
	// todo: add your logic here and delete this line

	return
}
