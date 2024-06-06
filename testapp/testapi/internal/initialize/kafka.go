package initialize

import (
	"dianshang/testapp/testapi/global"
	"dianshang/testapp/testapi/internal/dao/kafka"
	"github.com/IBM/sarama"
)

func SetupKafka() {
	config := global.Config.KafkaConfig
	global.KafkaBrokers = config.Brokers
	producer, err := sarama.NewSyncProducer(config.Brokers, nil)
	if err != nil {
		global.Logger.Fatal("Connect to Kafka failed: " + err.Error())
	}

	global.KafkaProducer = producer
	global.Logger.Info("Init Kafka success")
	// 创建 Topic 示例
	kafka.CreateTopic(global.KafkaBrokers, "my-topic", 1, 3)
}
