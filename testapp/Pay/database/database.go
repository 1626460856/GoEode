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

type MysqlUser struct { //用户注册信息
	ID           int     //自增键
	Username     string  // 用户名
	Password     string  // 用户密码
	Usernick     string  // 用户昵称
	UserIdentity string  // 用户身份唯一标识
	Balance      float64 // 余额
}

type Product struct {
	Id          int     `json:"id"`          // 哈希key键
	Name        string  `json:"name"`        // 商品名称
	Description string  `json:"description"` // 商品描述
	Price       float64 `json:"price"`       // 商品价格
	Stock       int     `json:"stock"`       // 商品库存
	Boss        string  `json:"boss"`        // 商品所属
}

func GetProduct(rdb *redis.Client, id int) (Product, error) {
	ctx := context.Background()
	// 尝试从 Redis 获取产品
	result, err := rdb.HGetAll(ctx, strconv.Itoa(id)).Result()
	if err != nil {
		// 如果 Redis 中没有产品，从 MySQL 获取
		product, err := GetProductById(ShopMySQLDB, id)
		if err != nil {
			return Product{}, err
		}

		// 将产品存储在 Redis 中，供未来的请求使用
		err = AddProductInRedis(ctx, rdb, product.Id, product.Name, product.Description, product.Price, product.Stock, product.Boss)
		if err != nil {
			return Product{}, err
		}

		return product, nil
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
func GetProductById(ShopMysqlDB *sql.DB, id int) (Product, error) {
	var product Product

	sqlStmt := `
	SELECT id, name, description, price, stock, boss
	FROM ProductList
	WHERE id = ?;`

	row := ShopMysqlDB.QueryRow(sqlStmt, id)
	err := row.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Stock, &product.Boss)
	if err != nil {
		if err == sql.ErrNoRows {
			// 没有找到对应的记录
			return Product{}, fmt.Errorf("没有找到ID为 %d 的商品", id)
		}

		// 其他错误
		return Product{}, fmt.Errorf("查询商品失败: %v", err)
	}

	return product, nil
}
func AddProductInRedis(ctx context.Context, ShopRedis1DB *redis.Client, id int, name string, description string, price float64, stock int, boss string) error {

	// 将商品信息存储在一个哈希中
	_, err := ShopRedis1DB.HSet(ctx, strconv.Itoa(id), map[string]interface{}{
		"name":        name,
		"description": description,
		"price":       price,
		"stock":       stock,
		"boss":        boss,
	}).Result()

	if err != nil {
		return fmt.Errorf("在 Redis 中创建商品失败: %v", err)
	}

	fmt.Println("在 Redis 中创建商品成功")
	return nil
}
