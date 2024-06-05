package initialize

import (
	"dianshang/testapp/testapi/global"
	"github.com/IBM/sarama"
)

func SetupKafka() {
	// 从全局配置中获取Kafka的配置信息
	config := global.Config.KafkaConfig

	// 配置Kafka生产者
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Producer.Retry.Max = 5
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll
	kafkaConfig.Producer.Partitioner = sarama.NewRandomPartitioner

	// 创建一个Kafka同步生产者
	kafkaProducer, err := sarama.NewSyncProducer(config.Brokers, kafkaConfig)
	if err != nil {
		global.Logger.Fatal("连接到Kafka失败: " + err.Error())
	}

	// 将Kafka生产者赋值给全局变量
	global.KafkaProducer = kafkaProducer

	// 记录初始化成功消息
	global.Logger.Info("初始化Kafka成功")
}
