package mysql

import (
	"context"
	"database/sql"
	"dianshang/testapp/testapi/global"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Id           int //自增键
	UserName     string
	PassWord     string
	UserNick     string
	UserIdentity string
	balance      float64 = 0
)

// CreateRegisterUsersTable 创建表格存储用户注册是数据
func CreateRegisterUsersTable() {
	// SQL语句
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS RegisterUsers (
		id INT AUTO_INCREMENT,
		username VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		usernick VARCHAR(255) NOT NULL,
		userIdentity VARCHAR(255) NOT NULL,
		balance FLOAT DEFAULT 0,
		PRIMARY KEY (id)
	);`

	// 执行SQL语句
	_, err := global.UserMysqlDB.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("创建RegisterUsers表失败: %v", err)
	}

	fmt.Println("RegisterUsers表创建成功")
}

// UpdateMysqlRegisterUsersFromRedis 更新 MySQL 数据库中的用户数据，使其与 Redis 中的数据保持一致
func UpdateMysqlRegisterUsersFromRedis(UserRedis1DB *redis.Client, UserMysqlDB *sql.DB) {
	ctx := context.Background()

	// 从 Redis 中获取所有用户的 username
	keys, err := UserRedis1DB.Keys(ctx, "*").Result()
	if err != nil {
		log.Fatalf("查询 Redis 失败: %v", err)
	}

	// 遍历所有 username
	for _, username := range keys {
		// 使用 TYPE 命令检查键的类型
		keyType, err := UserRedis1DB.Type(ctx, username).Result()
		if err != nil {
			log.Fatalf("获取 Redis 键的类型失败: %v", err)
		}

		// 如果键的类型是哈希，那么从哈希中获取用户数据
		if keyType == "hash" {
			result, err := UserRedis1DB.HGetAll(ctx, username).Result()
			if err != nil {
				log.Fatalf("从 Redis 获取用户失败: %v", err)
			}

			// 尝试插入用户数据到 MySQL 数据库中，如果用户已经存在，那么更新用户数据
			_, err = UserMysqlDB.Exec("INSERT INTO RegisterUsers (username, password, usernick, userIdentity, balance) VALUES (?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE password = ?, usernick = ?, userIdentity = ?, balance = ?",
				username, result["password"], result["usernick"], result["userIdentity"], result["balance"], result["password"], result["usernick"], result["userIdentity"], result["balance"])
			if err != nil {
				log.Fatalf("更新 MySQL 数据库失败: %v", err)
			}
		}
	}

	fmt.Println("成功从 Redis 更新到 MySQL")
}

// AddUserInMysql 新增或更新单个表格中单个数据
func AddUserInMysql(ctx context.Context, UserMysqlDB *sql.DB, username, password, usernick, userIdentity string, balance float64) error {
	// 尝试插入用户数据到 MySQL 数据库中，如果用户已经存在，那么更新用户数据
	_, err := UserMysqlDB.ExecContext(ctx, "INSERT INTO RegisterUsers (username, password, usernick, userIdentity, balance) VALUES (?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE password = ?, usernick = ?, userIdentity = ?, balance = ?",
		username, password, usernick, userIdentity, balance, password, usernick, userIdentity, balance)
	if err != nil {
		return fmt.Errorf("更新或新建 MySQL 数据库中的用户失败: %v", err)
	}

	fmt.Println("在 MySQL 中更新或新建用户成功")
	return nil
}
