package redis

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
)

type Product struct {
	Id          int     `json:"id"`          // 哈希key键
	Name        string  `json:"name"`        // 商品名称
	Description string  `json:"description"` // 商品描述
	Price       float64 `json:"price"`       // 商品价格
	Stock       int     `json:"stock"`       // 商品库存
	Boss        string  `json:"boss"`        // 商品所属
}

func AddProductInRedis(ctx context.Context, rdb *redis.Client, id int, name string, description string, price float64, stock int, boss string) error {

	// 将商品信息存储在一个哈希中
	_, err := rdb.HSet(ctx, strconv.Itoa(id), map[string]interface{}{
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

func GetProductById(ShopRedis2DB *redis.Client, id int) (Product, error) {
	ctx := context.Background()

	// 从哈希中获取商品信息
	result, err := ShopRedis2DB.HGetAll(ctx, strconv.Itoa(id)).Result()
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

// UpdateRedisProductListFromMysql 从 MySQL 更新 Redis 中的商品列表
func UpdateRedisProductListFromMysql(ShopMysqlDB *sql.DB, ShopRedis1DB *redis.Client) {
	ctx := context.Background()

	// 从 MySQL 数据库中查询所有商品数据
	rows, err := ShopMysqlDB.Query("SELECT id, name, description, price, stock, boss FROM ProductList")
	if err != nil {
		log.Fatalf("查询 MySQL 数据库失败: %v", err)
	}
	defer rows.Close()

	// 遍历查询结果
	for rows.Next() {
		var id int
		var name, description, boss string
		var price float64
		var stock int
		err = rows.Scan(&id, &name, &description, &price, &stock, &boss)
		if err != nil {
			log.Fatalf("读取 MySQL 数据库失败: %v", err)
		}

		// 将商品数据存储到 Redis 的哈希中
		err = AddProductInRedis(ctx, ShopRedis1DB, id, name, description, price, stock, boss)
		if err != nil {
			log.Fatalf("在 Redis 中存储商品失败: %v", err)
		}
	}

	fmt.Println("成功从 MySQL 更新到 Redis")
}
