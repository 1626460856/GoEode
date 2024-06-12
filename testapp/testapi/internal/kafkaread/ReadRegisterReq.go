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

type RegisterMessage struct { //注册kafka消息
	UserName     string `json:"UserName"`
	PassWord     string `json:"PassWord"`
	UserNick     string `json:"UserNick"`
	UserIdentity string `json:"UserIdentity"`
}

func ReadRegisterReq() { //读取注册kafka消息
	// 配置消费者
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        global.KafkaBrokers,
		Topic:          "RegisterReq",
		CommitInterval: 1 * time.Second,
		GroupID:        "group-id8",
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
		var msg RegisterMessage
		if err := json.Unmarshal(message.Value, &msg); err != nil {
			fmt.Printf("解码消息失败:%v\n", err)
			continue
		}
		mysql.AddUserInMysql(context.Background(), msg.UserName, msg.PassWord, msg.UserNick, msg.UserIdentity, 0)
		redis.AddUserInRedis(context.Background(), msg.UserName, msg.PassWord, msg.UserNick, msg.UserIdentity, 0)
		// 打印解码后的消息
		fmt.Printf("收到的信息 %s: %+v\n", "RegisterReq", msg)
		time.Sleep(1 * time.Second)
	}
}
