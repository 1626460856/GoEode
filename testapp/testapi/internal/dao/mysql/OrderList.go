package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
	"time"
)

type Order struct {
	OrderID     int     `json:"orderId"`     // 订单id  mysql自增键
	ProductID   int     `json:"productId"`   // 商品id
	ProductName string  `json:"productName"` // 商品名称
	Price       float64 `json:"price"`       // 商品价格
	Boss        string  `json:"boss"`        // 商家
	BuyQuantity int     `json:"buyQuantity"` // 购买商品数量
	UserName    string  `json:"userName"`    // 购买者这个通过传入的token解析获得
	Coupon      float64 `json:"coupon"`      // 优惠券,表格初始值默认为1
	OrderStatus string  `json:"orderStatus"` // 订单状态 有三种状态，“unpaid”为未支付，“paying”为支付中，“paid”为已支付，创建的时候默认未支付
	CreatedAt   string  `json:"createdAt"`   // 创建时间
	UpdatedAt   string  `json:"updatedAt"`   // 更新时间
}

func CreateOrderListTable(ShopMysqlDB *sql.DB) error {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS OrderList (
		orderId INT AUTO_INCREMENT,
		productId INT NOT NULL,
		productName VARCHAR(255) NOT NULL,
		price FLOAT NOT NULL,
		boss VARCHAR(255) NOT NULL,
		buyQuantity INT NOT NULL,
		userName VARCHAR(255) NOT NULL,
		coupon FLOAT DEFAULT 1,
		orderStatus VARCHAR(255) DEFAULT 'unpaid',
		createdAt VARCHAR(255),
		updatedAt VARCHAR(255),
		PRIMARY KEY (orderId)
	);`

	_, err := ShopMysqlDB.Exec(sqlStmt)
	if err != nil {
		return fmt.Errorf("创建OrderList表失败: %v", err)
	}

	fmt.Println("OrderList表创建成功")
	return nil
}

func AddOrderInMysql(ShopMysqlDB *sql.DB, ProductID int, ProductName string, Price float64, Boss string, BuyQuantity int, UserName string) (OrderID int, createdAt string, updatedAt string, err error) {
	// 获取当前时间
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	sqlStmt := `
	INSERT INTO OrderList (productId, productName, price, boss, buyQuantity, userName, coupon, orderStatus, createdAt, updatedAt) 
	VALUES (?, ?, ?, ?, ?, ?, 1, 'unpaid', ?, ?);`

	result, err := ShopMysqlDB.Exec(sqlStmt, ProductID, ProductName, Price, Boss, BuyQuantity, UserName, currentTime, currentTime)
	if err != nil {
		return 0, currentTime, currentTime, fmt.Errorf("向OrderList表插入数据失败: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, currentTime, currentTime, fmt.Errorf("获取插入的ID失败: %v", err)
	}

	fmt.Println("成功向OrderList表插入数据，OrderID为：", id)
	return int(id), currentTime, currentTime, nil
}

func GetOrderByIdInMysql(ShopMysqlDB *sql.DB, OrderID int) (Order, error) {
	var order Order

	sqlStmt := `
	SELECT orderId, productId, productName, price, boss, buyQuantity, userName, coupon, orderStatus, createdAt, updatedAt
	FROM OrderList
	WHERE orderId = ?;`

	row := ShopMysqlDB.QueryRow(sqlStmt, OrderID)
	err := row.Scan(&order.OrderID, &order.ProductID, &order.ProductName, &order.Price, &order.Boss, &order.BuyQuantity, &order.UserName, &order.Coupon, &order.OrderStatus, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			// 没有找到对应的记录
			return Order{}, fmt.Errorf("没有找到ID为 %d 的订单", OrderID)
		}

		// 其他错误
		return Order{}, fmt.Errorf("查询订单失败: %v", err)
	}

	return order, nil
}
func UpdateMysqlOrderListFromRedis(ShopRedis2DB *redis.Client, ShopMysqlDB *sql.DB) {
	ctx := context.Background()

	// 从 Redis 中获取所有订单的 ID
	keys, err := ShopRedis2DB.Keys(ctx, "*").Result()
	if err != nil {
		log.Fatalf("查询 Redis 失败: %v", err)
	}

	// 遍历所有 ID
	for _, id := range keys {
		// 使用 TYPE 命令检查键的类型
		keyType, err := ShopRedis2DB.Type(ctx, id).Result()
		if err != nil {
			log.Fatalf("获取 Redis 键的类型失败: %v", err)
		}

		// 如果键的类型是哈希，那么从哈希中获取订单数据
		if keyType == "hash" {
			result, err := ShopRedis2DB.HGetAll(ctx, id).Result()
			if err != nil {
				log.Fatalf("从 Redis 获取订单失败: %v", err)
			}

			// 将字符串转换为 float64 和 int
			orderID, _ := strconv.Atoi(id)
			productID, _ := strconv.Atoi(result["productId"])
			price, _ := strconv.ParseFloat(result["price"], 64)
			buyQuantity, _ := strconv.Atoi(result["buyQuantity"])
			coupon, _ := strconv.ParseFloat(result["coupon"], 64)

			// 尝试插入订单数据到 MySQL 数据库中，如果订单已经存在，那么更新订单数据
			_, err = ShopMysqlDB.Exec("INSERT INTO OrderList (orderId, productId, productName, price, boss, buyQuantity, userName, coupon, orderStatus, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE productId = ?, productName = ?, price = ?, boss = ?, buyQuantity = ?, userName = ?, coupon = ?, orderStatus = ?, createdAt = ?, updatedAt = ?",
				orderID, productID, result["productName"], price, result["boss"], buyQuantity, result["userName"], coupon, result["orderStatus"], result["createdAt"], result["updatedAt"], productID, result["productName"], price, result["boss"], buyQuantity, result["userName"], coupon, result["orderStatus"], result["createdAt"], result["updatedAt"])
			if err != nil {
				log.Fatalf("更新 MySQL 数据库失败: %v", err)
			}
		}
	}

	fmt.Println("成功从 Redis 更新到 MySQL")
}

// 删除一个订单
func DeleteOrderByIdInMysql(ShopMysqlDB *sql.DB, OrderID int) error {
	// 准备SQL语句
	stmt, err := ShopMysqlDB.Prepare("DELETE FROM OrderList WHERE OrderID = ?")
	if err != nil {
		return fmt.Errorf("SQL语句准备失败: %v", err)
	}
	defer stmt.Close()

	// 执行SQL语句
	_, err = stmt.Exec(OrderID)
	if err != nil {
		return fmt.Errorf("执行SQL语句失败: %v", err)
	}

	fmt.Println("从MySQL中成功删除订单")
	return nil
}
