package logic

import (
	"context"
	"dianshang/testapp/User/internal/svc"
	"dianshang/testapp/User/internal/types"
	"errors"
	"fmt"

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
	// 从JWT令牌中提取用户信息
	username, ok := l.ctx.Value("username").(string)
	if !ok {
		// 如果无法获取username，返回错误
		return nil, errors.New("username is missing from context")
	}
	fmt.Println("username:", username)

	// 使用提取到的用户信息
	// todo: add your logic here and delete this line
	return &types.UserInfoResply{
		Code:    200,
		Message: "成功获取用户信息",
		Data: &types.UserInfoItem{
			UserIdentity: "身份还未设置",
			UserName:     username,
			UserNick:     "查了再说",
		},
	}, nil
}
