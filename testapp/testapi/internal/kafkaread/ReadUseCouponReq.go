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

type UseCouponMessage struct {
	Coupon  float64 `json:"Coupon"`
	OrderId int     `json:"OrderId"`
}

func ReadUserCouponReq() { //读取创建订单kafka消息
	// 配置消费者
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        global.KafkaBrokers,
		Topic:          "UseCoupon",
		CommitInterval: 1 * time.Second,
		GroupID:        "ReadUseCouponReq",
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
		var msg UseCouponMessage
		if err := json.Unmarshal(message.Value, &msg); err != nil {
			fmt.Printf("解码消息失败:%v\n", err)
			continue
		}
		UpdatedAt, err := mysql.UseCouponInMysqlOrder(global.ShopMysqlDB, msg.OrderId, msg.Coupon)
		if err != nil {
			fmt.Printf("mysql更新订单失败:%v\n", err)
			continue

		}
		err = redis.UseCouponInRedisOrder(global.ShopRedis2DB, msg.OrderId, msg.Coupon, UpdatedAt)
		if err != nil {
			fmt.Printf("redis更新订单失败:%v\n", err)
			continue

		}
		// 打印解码后的消息
		fmt.Printf("收到的信息 %s: %+v\n", "UserCouponReq", msg)
		//time.Sleep(1 * time.Second)
	}
}
