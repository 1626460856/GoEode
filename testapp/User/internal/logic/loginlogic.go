package logic

import (
	"context"
	"dianshang/testapp/User/database"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"

	"dianshang/testapp/User/internal/svc"
	"dianshang/testapp/User/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.CommonResply, err error) {
	// todo: add your logic here and delete this line
	// 从Redis2中获取用户信息
	exists, err := database.UserRedis2DB.SIsMember(l.ctx, "UserName", req.UserName).Result()
	if err != nil {
		return nil, errors.New("Redis2查询失败")
	}
	if !exists {
		return nil, errors.New("用户不存在")
	}

	// 从Redis1中获取用户信息
	redisUser, err := database.GetRedisUserByUsername(database.UserRedis1DB, req.UserName)
	if err != nil || redisUser.Username == "" {
		// 如果Redis1中不存在，则从MySQL中获取
		mysqlUser, err := database.GetMysqlUserByUsername(database.UserMySQLDB, req.UserName)
		if err != nil {
			return nil, errors.New("在userredis2中发现但在userredis1中未发现，在mysql中也未发现")
		}
		// 将用户数据存储到 Redis1 的哈希中
		err = database.AddUserInRedis(l.ctx, mysqlUser.Username, mysqlUser.Password, mysqlUser.Usernick, mysqlUser.UserIdentity, mysqlUser.Balance)
		if err != nil {
			return nil, errors.New("Redis1更新失败")
		}
		redisUser = database.RedisUser{
			Username:     mysqlUser.Username,
			Password:     mysqlUser.Password,
			Usernick:     mysqlUser.Usernick,
			UserIdentity: mysqlUser.UserIdentity,
			Balance:      mysqlUser.Balance,
		}
	}

	// 验证密码
	if redisUser.Password != req.PassWord {
		return nil, errors.New("密码错误")
	}
	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.UserName,
		"exp":      time.Now().Add(time.Hour * 2).Unix(), //过期2h
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
		Message: "成功登录账户：" + req.UserName + "，并生成JWT Token",
		Data:    tokenString,
	}, nil
}
