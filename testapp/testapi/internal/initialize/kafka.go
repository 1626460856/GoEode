package initialize

import (
	"dianshang/testapp/testapi/global"
	"github.com/IBM/sarama"
)

func CreateTopic(brokers []string, topic string, partitions int32, replicationFactor int16) {
	admin, err := sarama.NewClusterAdmin(brokers, nil)
	if err != nil {
		global.Logger.Fatal("创建Kafka管理员失败: %v" + err.Error())
	}
	defer func() {
		if err := admin.Close(); err != nil {
			global.Logger.Error("关闭Kafka管理失败: %v" + err.Error())
		}
	}()

	// 检查 Topic 是否存在
	topics, err := admin.ListTopics()
	if err != nil {
		global.Logger.Fatal("无法列出主题: %v" + err.Error())
	}

	if _, exists := topics[topic]; !exists {
		// 创建 Topic
		err = admin.CreateTopic(topic, &sarama.TopicDetail{
			NumPartitions:     partitions,
			ReplicationFactor: replicationFactor,
		}, false)
		if err != nil {
			global.Logger.Fatal("创建主题" + topic + "失败" + err.Error())
		}
		global.Logger.Info(" 成功备份" + topic)
	} else {
		global.Logger.Info("消息主题" + topic + "在kafka集群已经存在")
	}
}

func SetupKafka() {
	// 创建用户主题名称
	brokers := global.Config.KafkaConfig.Brokers
	global.KafkaBrokers = brokers

	// 创建主题
	CreateTopic(global.KafkaBrokers, "RegisterReq", 5, 3)
	// 创建主题
	CreateTopic(global.KafkaBrokers, "AddProduct", 5, 3)
	// 创建主题
	CreateTopic(global.KafkaBrokers, "test2", 5, 3)

}
