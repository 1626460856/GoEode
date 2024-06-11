package redis

import (
	"context"
	"dianshang/testapp/testapi/global"
	"fmt"
	"log"
)

// AddUserInRedis 新增或更新单个数据的哈希值
func AddUserInRedis(ctx context.Context, username, password, usernick, userIdentity string, balance float64) error {
	err := global.UserRedis1DB.HSet(ctx, username, map[string]interface{}{
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

// UpdateRedisFromMysql 更新 Redis 中的用户数据，使其与 MySQL 数据库中的数据保持一致
func UpdateRedisFromMysql() {
	ctx := context.Background()

	// 从 MySQL 数据库中查询所有用户数据
	rows, err := global.UserMysqlDB.Query("SELECT username, password, usernick, userIdentity, balance FROM RegisterUsers")
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
		err = AddUserInRedis(ctx, username, password, usernick, userIdentity, balance)
		if err != nil {
			log.Fatalf("在 Redis 中存储用户失败: %v", err)
		}
	}

	fmt.Println("成功从 MySQL 更新到 Redis")
}
