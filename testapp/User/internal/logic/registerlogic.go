package logic

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"time"

	"dianshang/testapp/User/internal/svc"
	"dianshang/testapp/User/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.CommonResply, error) {
	// todo: add your logic here and delete this line

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.UserName,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	// 使用secret签名并获得完整的编码令牌作为字符串
	tokenString, err := token.SignedString([]byte("Eode"))
	if err != nil {
		return nil, err
	}
	// 将令牌保存到数据库或其他持久存储中

	// 返回令牌和用户名
	return &types.CommonResply{
		Code:    200,
		Message: "成功创建账户：" + req.UserName + "，并生成JWT Token",
		Data:    tokenString,
	}, nil
}
