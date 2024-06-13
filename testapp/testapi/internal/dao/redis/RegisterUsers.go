package redis

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
)

type UserInRedis struct {
	UserName     string //key键
	PassWord     string
	UserNick     string
	UserIdentity string
	balance      float64 //= 0
}

// AddUserInRedis 新增或更新单个数据的哈希值
func AddUserInRedis(ctx context.Context, UserRedis1DB *redis.Client, username, password, usernick, userIdentity string, balance float64) error {
	err := UserRedis1DB.HSet(ctx, username, map[string]interface{}{
		"password":     password,
		"usernick":     usernick,
		"userIdentity": userIdentity,
		"balance":      balance,
	}).Err()
	if err != nil {
		return fmt.Errorf("在Redis中存储用户失败: %v", err)
	}

	fmt.Println("在Redis中存储用户成功")
	return nil
}

// UpdateRedisRegisterUsersFromMysql 更新 Redis 中的用户数据，使其与 MySQL 数据库中的数据保持一致
func UpdateRedisRegisterUsersFromMysql(UserRedis1DB *redis.Client, UserMysqlDB *sql.DB) {
	ctx := context.Background()

	// 从 MySQL 数据库中查询所有用户数据
	rows, err := UserMysqlDB.Query("SELECT username, password, usernick, userIdentity, balance FROM RegisterUsers")
	if err != nil {
		log.Fatalf("查询 MySQL 数据库失败: %v", err)
	}
	defer rows.Close()

	// 遍历查询结果
	for rows.Next() {
		var username, password, usernick, userIdentity string
		var balance float64
		err = rows.Scan(&username, &password, &usernick, &userIdentity, &balance)
		if err != nil {
			log.Fatalf("读取 MySQL 数据库失败: %v", err)
		}

		// 将用户数据存储到 Redis 的哈希中
		err = AddUserInRedis(ctx, UserRedis1DB, username, password, usernick, userIdentity, balance)
		if err != nil {
			log.Fatalf("在 Redis 中存储用户失败: %v", err)
		}
	}

	fmt.Println("成功从 MySQL 更新到 Redis")
}
func GetUserByUsernameInRedis(UserRedis1DB *redis.Client, username string) (UserInRedis, error) {
	ctx := context.Background()

	// 从 Redis 中获取用户数据
	result, err := UserRedis1DB.HGetAll(ctx, username).Result()
	if err != nil {
		return UserInRedis{}, fmt.Errorf("从 Redis 获取用户失败: %v", err)
	}

	// 将 balance 转换为 float64
	balance, err := strconv.ParseFloat(result["balance"], 64)
	if err != nil {
		return UserInRedis{}, fmt.Errorf("转换 balance 失败: %v", err)
	}

	// 创建一个 UserInRedis 结构体
	user := UserInRedis{
		UserName:     username,
		PassWord:     result["password"],
		UserNick:     result["usernick"],
		UserIdentity: result["userIdentity"],
		balance:      balance,
	}

	return user, nil
}

// ChangeUserBalanceInRedis 更改 Redis 中用户的余额
func ChangeUserBalanceInRedis(UserRedis1DB *redis.Client, username string, changebalance float64) error {
	ctx := context.Background()

	// 获取当前余额
	currentBalanceStr, err := UserRedis1DB.HGet(ctx, username, "balance").Result()
	if err != nil {
		return fmt.Errorf("在 Redis 中获取用户余额失败: %v", err)
	}

	// 将余额从字符串转换为 float64
	currentBalance, err := strconv.ParseFloat(currentBalanceStr, 64)
	if err != nil {
		return fmt.Errorf("转换余额失败: %v", err)
	}

	// 计算新的余额
	newBalance := currentBalance + changebalance

	// 更新余额
	err = UserRedis1DB.HSet(ctx, username, "balance", newBalance).Err()
	if err != nil {
		return fmt.Errorf("在 Redis 中更新用户余额失败: %v", err)
	}

	fmt.Println("在 Redis 中更新用户余额成功")
	return nil
}
