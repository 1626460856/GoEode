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

type SeckillRequestMessage struct {
	ProductID   int    `json:"productId"`
	BuyQuantity int    `json:"buyQuantity"`
	UserName    string `json:"userName"`
}

func ReadSeckillRequestReq() { //读取创建订单kafka消息
	// 配置消费者
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        global.KafkaBrokers,
		Topic:          "SeckillRequest",
		CommitInterval: 1 * time.Second,
		GroupID:        "ReadSeckillRequestReq",
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
		var msg SeckillRequestMessage
		if err := json.Unmarshal(message.Value, &msg); err != nil {
			fmt.Printf("解码消息失败:%v\n", err)
			continue
		}
		// 打印解码后的消息
		fmt.Printf("收到的信息 %s: %+v\n", "PayOrderReq", msg)

		//获取商品信息
		product, err := redis.GetProductById(global.ShopRedis1DB, msg.ProductID)
		if err != nil {
			fmt.Printf("redis查询订单失败:%v\n", err)
			continue
		}
		//更新mysql库存
		err = mysql.DecreaseProductStock(global.ShopMysqlDB, product.Id, msg.BuyQuantity)
		if err != nil {
			fmt.Printf("mysql更新商品库存失败:%v\n", err)
			continue
		}

		//更新mysql和redis买家余额
		err = mysql.ChangeUserBalanceInMysql(global.UserMysqlDB, msg.UserName, -(product.Price * float64(msg.BuyQuantity)))
		if err != nil {
			fmt.Printf("mysql更新买家余额失败:%v\n", err)
			continue
		}
		err = mysql.ChangeUserBalanceInMysql(global.UserMysqlDB, product.Boss, product.Price*float64(msg.BuyQuantity))
		if err != nil {
			fmt.Printf("mysql更新商家余额失败:%v\n", err)
			continue
		}
		//更新mysql和redis商家余额
		err = redis.ChangeUserBalanceInRedis(global.UserRedis1DB, msg.UserName, -(product.Price * float64(msg.BuyQuantity)))
		if err != nil {
			fmt.Printf("redis更新买家余额失败:%v\n", err)
			continue
		}
		err = redis.ChangeUserBalanceInRedis(global.UserRedis1DB, product.Boss, product.Price*float64(msg.BuyQuantity))
		if err != nil {
			fmt.Printf("redis更新商家余额失败:%v\n", err)
			continue
		}
		//创建订单
		orderId, createdAt, updatedAt, err := mysql.AddSeckillOrderInMysql(global.ShopMysqlDB, product.Id, product.Name, product.Price, product.Boss, msg.BuyQuantity, msg.UserName)
		if err != nil {
			fmt.Printf("mysql创建秒杀订单失败:%v\n", err)
			continue
		}
		err = redis.AddOrderInRedis(context.Background(), global.ShopRedis2DB, orderId, product.Id, product.Name, product.Price, product.Boss, msg.BuyQuantity, msg.UserName, 1, "Seckillpaid", createdAt, updatedAt)
		if err != nil {
			fmt.Printf("redis创建秒杀订单失败:%v\n", err)
			continue
		}
		//time.Sleep(1 * time.Second)

	}
}
