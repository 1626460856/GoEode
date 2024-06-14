package kafkaread

import (
	"context"
	"dianshang/testapp/testapi/global"
	"dianshang/testapp/testapi/internal/dao/mysql"
	"dianshang/testapp/testapi/internal/dao/redis"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/segmentio/kafka-go"
	"time"
)

type AddProductMessage struct { //添加商品kafka消息
	Name        string  `json:"Name"`        // 商品名称
	Description string  `json:"Description"` // 商品描述
	Price       float64 `json:"Price"`       // 商品价格
	Stock       int     `json:"Stock"`       // 商品库存
	Boss        string  `json:"Boss"`        // 商品所属
}

func ReadAddProductReq() { //读取添加商品kafka消息
	// 配置消费者
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        global.KafkaBrokers,
		Topic:          "AddProduct",
		CommitInterval: 1 * time.Second,
		GroupID:        "ReadAddProductReq",
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
		var msg AddProductMessage
		if err := json.Unmarshal(message.Value, &msg); err != nil {
			fmt.Printf("解码消息失败:%v\n", err)
			continue
		}
		id, _ := mysql.AddProductInMysql(global.ShopMysqlDB, msg.Name, msg.Description, msg.Price, msg.Stock, msg.Boss)
		redis.AddProductInRedis(context.Background(), global.ShopRedis1DB, id, msg.Name, msg.Description, msg.Price, msg.Stock, msg.Boss)
		// 打印解码后的消息
		fmt.Printf("收到的信息 %s: %+v\n", "AddProductReq", msg)
		//time.Sleep(1 * time.Second)
	}
}
