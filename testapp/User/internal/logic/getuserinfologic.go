package logic

import (
	"context"
	"dianshang/testapp/User/database"
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
	// 从Redis2中获取用户信息
	exists, err := database.UserRedis2DB.SIsMember(l.ctx, "UserName", username).Result()
	if err != nil {
		return nil, errors.New("Redis2查询失败")
	}
	if !exists {
		return nil, errors.New("用户不存在")
	}

	// 从Redis1中获取用户信息
	redisUser, err := database.GetRedisUserByUsername(database.UserRedis1DB, username)
	if err != nil || redisUser.Username == "" {
		// 如果Redis1中不存在，则从MySQL中获取
		mysqlUser, err := database.GetMysqlUserByUsername(database.UserMySQLDB, username)
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
	// 使用提取到的用户信息
	// todo: add your logic here and delete this line
	return &types.UserInfoResply{
		Code:    200,
		Message: "成功获取用户信息",
		Data: &types.UserInfoItem{
			UserIdentity: redisUser.UserIdentity,
			UserName:     redisUser.Username,
			UserNick:     redisUser.Usernick,
			Balance:      redisUser.Balance,
			PassWord:     redisUser.Password,
		},
	}, nil
}
