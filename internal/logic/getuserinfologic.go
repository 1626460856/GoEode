package logic

import (
	"context"

	"dianshang/internal/svc"
	"dianshang/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetuserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetuserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetuserInfoLogic {
	return &GetuserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetuserInfoLogic) GetuserInfo(req *types.UserInfoReq) (resp *types.UserInfoResply, err error) {
	// todo: add your logic here and delete this line

	return
}
