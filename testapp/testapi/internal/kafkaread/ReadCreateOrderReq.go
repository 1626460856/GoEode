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

type CreateOrderMessage struct { //发送kafka消息
	// 订单id  mysql自增键
	ProductID int `json:"productId"` // 商品id传入获得
	// 商品名称kafka之后查询
	//商品价格kafka之后查询
	// 商家kafka之后查询
	BuyQuantity int    `json:"buyQuantity"` // 购买商品数量
	UserName    string `json:"userName"`    // 购买者这个通过传入的token解析获得
	// 优惠券,创建时默认为1
	// 订单状态 有三种状态，“unpaid”为未支付，“paying”为支付中，“paid”为已支付，创建的时候默认未支付
	// 创建时间kafka之后
	// 更新时间kafka之后
}

func ReadCreateOrderReq() { //读取创建订单kafka消息
	// 配置消费者
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        global.KafkaBrokers,
		Topic:          "CreateOrder",
		CommitInterval: 1 * time.Second,
		GroupID:        "ReadCreateOrderReq",
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
		var msg CreateOrderMessage
		if err := json.Unmarshal(message.Value, &msg); err != nil {
			fmt.Printf("解码消息失败:%v\n", err)
			continue
		}
		product, err := redis.GetProductById(global.ShopRedis1DB, msg.ProductID)
		if err != nil {
			fmt.Printf("根据商品id获取将购买商品失败: %v\n", err)
			continue
		}
		id, createdAt, updatedAt, err := mysql.AddOrderInMysql(global.ShopMysqlDB, product.Id, product.Name, product.Price, product.Boss, msg.BuyQuantity, msg.UserName)
		if err != nil {
			fmt.Printf("mysql创建订单失败: %v\n", err)
			continue
		}
		err = redis.AddOrderInRedis(context.Background(), global.ShopRedis2DB, id, product.Id, product.Name, product.Price, product.Boss, msg.BuyQuantity, msg.UserName, 1, "unpaid", createdAt, updatedAt)
		if err != nil {
			fmt.Printf("redis创建订单失败: %v\n", err)
			continue
		
		}
		// 打印解码后的消息
		fmt.Printf("收到的信息 %s: %+v\n", "CreateOrderReq", msg)
		time.Sleep(1 * time.Second)
	}
}
