package kafkaread

import (
	"context"
	"dianshang/testapp/testapi/global"
	"dianshang/testapp/testapi/internal/dao/mysql"
	"dianshang/testapp/testapi/internal/dao/redis"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

type PayOrderMessage struct {
	OrderId     int     `json:"OrderId"`
	Username    string  `json:"Username"`
	Boss        string  `json:"Boss"`
	Price       float64 `json:"Price"`
	BuyQuantity int     `json:"BuyQuantity"`
	Coupon      float64 `json:"Coupon"`
}

func ReadPayOrderReq() { //读取创建订单kafka消息
	// 配置消费者
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        global.KafkaBrokers,
		Topic:          "PayOrder",
		CommitInterval: 1 * time.Second,
		GroupID:        "ReadPayOrderReq",
		StartOffset:    kafka.FirstOffset,
	})

	ctx := context.Background()

	// 死循环一直读取消息
	for {
		message, err := reader.ReadMessage(ctx)
		if err != nil {
			fmt.Printf("读取kafka失败:%v\n", err)
			break
		}

		// 解码消息
		var msg PayOrderMessage
		if err := json.Unmarshal(message.Value, &msg); err != nil {
			fmt.Printf("解码消息失败:%v\n", err)
			continue
		}
		// 打印解码后的消息
		fmt.Printf("收到的信息 %s: %+v\n", "PayOrderReq", msg)

		//获取订单信息
		order, err := redis.GetOrderById(global.ShopRedis2DB, msg.OrderId)
		if err != nil {
			fmt.Printf("redis查询订单失败:%v\n", err)
			continue
		}
		//更新mysql和redis库存
		err = mysql.DecreaseProductStock(global.ShopMysqlDB, order.ProductID, order.BuyQuantity)
		if err != nil {
			fmt.Printf("mysql更新商品库存失败:%v\n", err)
			continue
		}
		err = redis.DecreaseProductStock(global.ShopRedis1DB, order.ProductID, order.BuyQuantity)
		if err != nil {
			fmt.Printf("redis更新商品库存失败:%v\n", err)
			continue
		}
		//更新mysql和redis买家余额
		err = mysql.ChangeUserBalanceInMysql(global.UserMysqlDB, msg.Username, -(msg.Price * float64(msg.BuyQuantity) * msg.Coupon))
		if err != nil {
			fmt.Printf("mysql更新买家余额失败:%v\n", err)
			continue
		}
		err = mysql.ChangeUserBalanceInMysql(global.UserMysqlDB, msg.Boss, msg.Price*float64(msg.BuyQuantity)*msg.Coupon)
		if err != nil {
			fmt.Printf("mysql更新商家余额失败:%v\n", err)
			continue
		}
		//更新mysql和redis商家余额
		err = redis.ChangeUserBalanceInRedis(global.UserRedis1DB, msg.Username, -(msg.Price * float64(msg.BuyQuantity) * msg.Coupon))
		if err != nil {
			fmt.Printf("redis更新买家余额失败:%v\n", err)
			continue
		}
		err = redis.ChangeUserBalanceInRedis(global.UserRedis1DB, msg.Boss, msg.Price*float64(msg.BuyQuantity)*msg.Coupon)
		if err != nil {
			fmt.Printf("redis更新商家余额失败:%v\n", err)
			continue
		}
		UpdatedAt, err := mysql.UpdateOrderToPaidInMysql(global.ShopMysqlDB, msg.OrderId)
		if err != nil {
			fmt.Printf("mysql更新订单状态失败:%v\n", err)
			continue

		}
		err = redis.UpdateOrderToPaidInRedis(global.ShopRedis2DB, msg.OrderId, UpdatedAt)
		if err != nil {
			fmt.Printf("redis更新订单状态失败:%v\n", err)
			continue

		}
		time.Sleep(1 * time.Second)

	}
}
