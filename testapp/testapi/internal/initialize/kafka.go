package initialize

import (
	"dianshang/testapp/testapi/global"
	"dianshang/testapp/testapi/internal/dao/kafka"
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

func SetupKafka() {
	// 主题名称
	topic := "testtopic"
	brokers := global.Config.KafkaConfig.Brokers
	global.KafkaBrokers = brokers

	// 创建主题
	kafka.CreateTopic(global.KafkaBrokers, topic, 2, 3)

	// 配置消费者
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRange}
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumer, err := sarama.NewConsumer(global.KafkaBrokers, config)
	if err != nil {
		log.Fatalf("启动Kafka消费者失败: %v", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("启动分区消费者失败: %v", err)
	}
	defer partitionConsumer.Close()

	go func() {
		for message := range partitionConsumer.Messages() {
			// 处理消息
			fmt.Printf("Received message: %s\n", string(message.Value))
			// 在这里处理用户注册逻辑，例如与数据库交互等
		}
	}()
}
