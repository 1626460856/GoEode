package kafkaread

import (
	"context"
	"dianshang/testapp/testapi/global"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

func ReadTest2Req() { //读取注册kafka消息
	// 配置消费者
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        global.KafkaBrokers,
		Topic:          "test2",
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
			continue
		}

		// 打印解码后的消息
		fmt.Printf("收到的信息 %s: %+v\n", "test2", message)
		time.Sleep(1 * time.Second)
	}
}
