package redis

import (
	"context"
	"database/sql"
	"dianshang/testapp/testapi/global"
	"dianshang/testapp/testapi/internal/dao/mysql"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
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

func AddOrderInRedis(ctx context.Context, ShopRedis2DB *redis.Client, OrderID int, ProductID int, ProductName string, Price float64, Boss string, BuyQuantity int, UserName string, Coupon float64, OrderStatus string, CreatedAt string, UpdatedAt string) error {
	// 将订单信息存储在一个哈希中
	_, err := ShopRedis2DB.HSet(ctx, strconv.Itoa(OrderID), map[string]interface{}{
		"productId":   ProductID,
		"productName": ProductName,
		"price":       Price,
		"boss":        Boss,
		"buyQuantity": BuyQuantity,
		"userName":    UserName,
		"coupon":      Coupon,
		"orderStatus": OrderStatus,
		"createdAt":   CreatedAt,
		"updatedAt":   UpdatedAt,
	}).Result()

	if err != nil {
		return fmt.Errorf("在 Redis 中创建订单失败: %v", err)
	}

	fmt.Println("在 Redis 中创建订单成功")
	return nil
}

func GetOrderByIdInRedis(ctx context.Context, ShopRedis2DB *redis.Client, OrderID int) (Order, error) {
	var order Order

	// 从Redis中获取订单信息
	result, err := ShopRedis2DB.HGetAll(ctx, strconv.Itoa(OrderID)).Result()
	if err != nil {
		return Order{}, fmt.Errorf("在 Redis 中查询订单失败: %v", err)
	}

	// 将获取的数据赋值给 Order 结构体的字段
	order.OrderID = OrderID
	order.ProductID, _ = strconv.Atoi(result["productId"])
	order.ProductName = result["productName"]
	order.Price, _ = strconv.ParseFloat(result["price"], 64)
	order.Boss = result["boss"]
	order.BuyQuantity, _ = strconv.Atoi(result["buyQuantity"])
	order.UserName = result["userName"]
	order.Coupon, _ = strconv.ParseFloat(result["coupon"], 64)
	order.OrderStatus = result["orderStatus"]
	order.CreatedAt = result["createdAt"]
	order.UpdatedAt = result["updatedAt"]

	return order, nil
}

// 先在redis找再在mysql找
func GetOrderById(ShopRedis2DB *redis.Client, OrderId int) (Order, error) {
	ctx := context.Background()
	// 尝试从 Redis 获取产品
	result, err := ShopRedis2DB.HGetAll(ctx, strconv.Itoa(OrderId)).Result()
	if err != nil {
		// 如果 Redis 中没有产品，从 MySQL 获取
		mysqlOrder, err := mysql.GetOrderByIdInMysql(global.ShopMysqlDB, OrderId)
		if err != nil {
			return Order{}, err
		}

		// 将产品存储在 Redis 中，供未来的请求使用
		err = AddOrderInRedis(ctx, ShopRedis2DB, mysqlOrder.OrderID, mysqlOrder.ProductID, mysqlOrder.ProductName, mysqlOrder.Price, mysqlOrder.Boss, mysqlOrder.BuyQuantity, mysqlOrder.UserName, mysqlOrder.Coupon, mysqlOrder.OrderStatus, mysqlOrder.CreatedAt, mysqlOrder.UpdatedAt)
		if err != nil {
			return Order{}, err
		}
		order := Order{
			OrderID:     mysqlOrder.OrderID,
			ProductID:   mysqlOrder.ProductID,
			ProductName: mysqlOrder.ProductName,
			Price:       mysqlOrder.Price,
			Boss:        mysqlOrder.Boss,
			BuyQuantity: mysqlOrder.BuyQuantity,
			UserName:    mysqlOrder.UserName,
			Coupon:      mysqlOrder.Coupon,
			OrderStatus: mysqlOrder.OrderStatus,
			CreatedAt:   mysqlOrder.CreatedAt,
			UpdatedAt:   mysqlOrder.UpdatedAt,
		}
		return order, nil
	}
	var order Order
	// 将获取的数据赋值给 Order 结构体的字段
	order.OrderID = OrderId
	order.ProductID, _ = strconv.Atoi(result["productId"])
	order.ProductName = result["productName"]
	order.Price, _ = strconv.ParseFloat(result["price"], 64)
	order.Boss = result["boss"]
	order.BuyQuantity, _ = strconv.Atoi(result["buyQuantity"])
	order.UserName = result["userName"]
	order.Coupon, _ = strconv.ParseFloat(result["coupon"], 64)
	order.OrderStatus = result["orderStatus"]
	order.CreatedAt = result["createdAt"]
	order.UpdatedAt = result["updatedAt"]

	return order, nil
}

func UpdateRedisOderListFromMysql(ShopMysqlDB *sql.DB, ShopRedis2DB *redis.Client) {
	ctx := context.Background()

	// 从 MySQL 数据库中查询所有订单数据
	rows, err := ShopMysqlDB.Query("SELECT OrderID, ProductID, ProductName, Price, Boss, BuyQuantity, UserName, Coupon, OrderStatus, CreatedAt, UpdatedAt FROM OrderList")
	if err != nil {
		log.Fatalf("查询 MySQL 数据库失败: %v", err)
	}
	defer rows.Close()

	// 遍历查询结果
	for rows.Next() {
		var OrderID, ProductID, BuyQuantity int
		var ProductName, Boss, UserName, OrderStatus, CreatedAt, UpdatedAt string
		var Price, Coupon float64
		err = rows.Scan(&OrderID, &ProductID, &ProductName, &Price, &Boss, &BuyQuantity, &UserName, &Coupon, &OrderStatus, &CreatedAt, &UpdatedAt)
		if err != nil {
			log.Fatalf("读取 MySQL 数据库失败: %v", err)
		}

		// 调用 AddOrderInRedis 函数将订单数据存储到 Redis 的哈希中
		err = AddOrderInRedis(ctx, ShopRedis2DB, OrderID, ProductID, ProductName, Price, Boss, BuyQuantity, UserName, Coupon, OrderStatus, CreatedAt, UpdatedAt)
		if err != nil {
			log.Fatalf("在 Redis 中存储订单失败: %v", err)
		}
	}

	fmt.Println("成功从 MySQL 更新到 Redis")
}
func DeleteOrderByIdInRedis(ShopRedis2DB *redis.Client, OrderID int) error {
	ctx := context.Background()
	// 从Redis中删除订单
	err := ShopRedis2DB.Del(ctx, strconv.Itoa(OrderID)).Err()
	if err != nil {
		return fmt.Errorf("删除 Redis 中的订单失败: %v", err)
	}

	fmt.Println("在 Redis 中成功删除订单")
	return nil
}
