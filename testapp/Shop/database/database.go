package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

var UserRedis2DB = redis.NewClient(&redis.Options{
	Addr:     "localhost:26379",
	Password: "123awzsex",
	DB:       0, // use default DB
})
var UserRedis1DB = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "123awzsex",
	DB:       0, // use default DB
})
var ShopRedis1DB = redis.NewClient(&redis.Options{
	Addr:     "localhost:6380",
	Password: "root",
	DB:       0, // use default DB
})
var ShopRedis2DB = redis.NewClient(&redis.Options{
	Addr:     "localhost:26380",
	Password: "root",
	DB:       0, // use default DB
})
var ShopMySQLDB *sql.DB

func init() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s",
		"root",
		"root",
		"localhost",
		3307,
		"test",
		"utf8mb4",
		"Asia%2FShanghai",
	)
	ShopMySQLDB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	ShopMySQLDB.SetConnMaxIdleTime(1 * time.Hour)
	ShopMySQLDB.SetConnMaxLifetime(1 * time.Hour)
	ShopMySQLDB.SetMaxIdleConns(10)
	ShopMySQLDB.SetMaxOpenConns(5)
}

// AddUserInRedis 新增或更新单个数据的哈希值
func AddUserInRedis(ctx context.Context, username, password, usernick, userIdentity string, balance float64) error {
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

type MysqlUser struct { //用户注册信息
	ID           int     //自增键
	Username     string  // 用户名
	Password     string  // 用户密码
	Usernick     string  // 用户昵称
	UserIdentity string  // 用户身份唯一标识
	Balance      float64 // 余额
}

func GetMysqlUserByUsername(MySQLDB *sql.DB, username string) (MysqlUser, error) {
	var user MysqlUser
	query := "SELECT id, username, password, usernick, userIdentity, balance FROM RegisterUsers WHERE username = ?"
	err := MySQLDB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Usernick, &user.UserIdentity, &user.Balance)
	if err != nil {
		return MysqlUser{}, fmt.Errorf("查询 MySQL 数据库失败: %v", err)
	}
	return user, nil
}

type RedisUser struct { //用户注册信息
	Username     string  // 用户名作为了哈希键
	Password     string  // 用户密码
	Usernick     string  // 用户昵称
	UserIdentity string  // 用户身份唯一标识
	Balance      float64 // 余额
}

func GetRedisUserByUsername(RedisDB *redis.Client, username string) (RedisUser, error) {
	var user RedisUser
	ctx := context.Background()

	result, err := RedisDB.HGetAll(ctx, username).Result()
	if err != nil {
		return RedisUser{}, fmt.Errorf("查询 Redis 失败: %v", err)
	}

	user.Username = username
	user.Password = result["password"]
	user.Usernick = result["usernick"]
	user.UserIdentity = result["userIdentity"]

	// 将余额从字符串转换为浮点数
	balance, err := strconv.ParseFloat(result["balance"], 64)
	if err != nil {
		return RedisUser{}, fmt.Errorf("转换余额失败: %v", err)
	}
	user.Balance = balance

	return user, nil
}

type Product struct {
	Id          int     `json:"id"`          // 哈希key键
	Name        string  `json:"name"`        // 商品名称
	Description string  `json:"description"` // 商品描述
	Price       float64 `json:"price"`       // 商品价格
	Stock       int     `json:"stock"`       // 商品库存
	Boss        string  `json:"boss"`        // 商品所属
}

func GetProductById(rdb *redis.Client, id int) (Product, error) {
	ctx := context.Background()

	// 从哈希中获取商品信息
	result, err := rdb.HGetAll(ctx, strconv.Itoa(id)).Result()
	if err != nil {
		return Product{}, fmt.Errorf("从 Redis 获取商品失败: %v", err)
	}

	// 将获取的数据赋值给 Product 结构体的字段
	product := Product{
		Id:          id,
		Name:        result["name"],
		Description: result["description"],
	}

	// 将字符串转换为 float64
	product.Price, err = strconv.ParseFloat(result["price"], 64)
	if err != nil {
		return Product{}, fmt.Errorf("解析价格失败: %v", err)
	}

	// 将字符串转换为 int
	product.Stock, err = strconv.Atoi(result["stock"])
	if err != nil {
		return Product{}, fmt.Errorf("解析库存失败: %v", err)
	}

	product.Boss = result["boss"]

	return product, nil
}
