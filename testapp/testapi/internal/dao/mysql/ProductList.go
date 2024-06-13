package mysql

import (
	"context"
	"database/sql"
	"dianshang/testapp/testapi/global"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
)

type Product struct {
	Id          int     `json:"id"`          // 商品ID根据kafka消息交互mysql自增键生成
	Name        string  `json:"name"`        // 商品名称
	Description string  `json:"description"` // 商品描述
	Price       float64 `json:"price"`       // 商品价格
	Stock       int     `json:"stock"`       // 商品库存
	Boss        string  `json:"boss"`        // 商品所属
}

// CreateProductTable使用Product结构在MySQL数据库中创建一个表
func CreateProductListTable() {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS ProductList (
		id INT AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,
		description VARCHAR(255) NOT NULL,
		price FLOAT NOT NULL,
		stock INT NOT NULL,
		boss VARCHAR(255) NOT NULL,
		PRIMARY KEY (id)
	);`

	_, err := global.ShopMysqlDB.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("创建Product表失败: %v", err)
	}

	fmt.Println("ProductList表创建成功")
}

func AddProductInMysql(db *sql.DB, name string, description string, price float64, stock int, boss string) (ID int, err error) {
	sqlStmt := `
	INSERT INTO ProductList (name, description, price, stock, boss) 
	VALUES (?, ?, ?, ?, ?);`

	result, err := db.Exec(sqlStmt, name, description, price, stock, boss)
	if err != nil {
		return 0, fmt.Errorf("向ProductList表插入数据失败: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("获取插入的ID失败: %v", err)
	}

	fmt.Println("成功向ProductList表插入数据，ID为：", id)
	return int(id), nil
}
func GetProductByIdInMysql(db *sql.DB, id int) (Product, error) {
	var product Product

	sqlStmt := `
	SELECT id, name, description, price, stock, boss
	FROM ProductList
	WHERE id = ?;`

	row := db.QueryRow(sqlStmt, id)
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

// UpdateMysqlProductListFromRedis 从 Redis 更新 MySQL 数据库中的商品列表
func UpdateMysqlProductListFromRedis(ShopRedis1DB *redis.Client, ShopMysqlDB *sql.DB) {
	ctx := context.Background()

	// 从 Redis 中获取所有商品的 ID
	keys, err := ShopRedis1DB.Keys(ctx, "*").Result()
	if err != nil {
		log.Fatalf("查询 Redis 失败: %v", err)
	}

	// 遍历所有 ID
	for _, id := range keys {
		// 使用 TYPE 命令检查键的类型
		keyType, err := ShopRedis1DB.Type(ctx, id).Result()
		if err != nil {
			log.Fatalf("获取 Redis 键的类型失败: %v", err)
		}

		// 如果键的类型是哈希，那么从哈希中获取商品数据
		if keyType == "hash" {
			result, err := ShopRedis1DB.HGetAll(ctx, id).Result()
			if err != nil {
				log.Fatalf("从 Redis 获取商品失败: %v", err)
			}

			// 将字符串转换为 float64 和 int
			price, _ := strconv.ParseFloat(result["price"], 64)
			stock, _ := strconv.Atoi(result["stock"])

			// 尝试插入商品数据到 MySQL 数据库中，如果商品已经存在，那么更新商品数据
			_, err = ShopMysqlDB.Exec("INSERT INTO ProductList (id, name, description, price, stock, boss) VALUES (?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE name = ?, description = ?, price = ?, stock = ?, boss = ?",
				id, result["name"], result["description"], price, stock, result["boss"], result["name"], result["description"], price, stock, result["boss"])
			if err != nil {
				log.Fatalf("更新 MySQL 数据库失败: %v", err)
			}
		}
	}

	fmt.Println("成功从 Redis 更新到 MySQL")
}
