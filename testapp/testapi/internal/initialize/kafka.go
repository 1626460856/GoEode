package initialize

import (
	"dianshang/testapp/testapi/global"
	"dianshang/testapp/testapi/internal/dao/kafka"
)

func SetupKafka() {

	// 主题名称
	topic := "testtopic"
	config := global.Config.KafkaConfig.Brokers
	global.KafkaBrokers = config

	//fmt.Println(global.KafkaBrokers)
	kafka.CreateTopic(global.KafkaBrokers, topic, 2, 3)

}
