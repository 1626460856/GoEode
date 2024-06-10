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
			global.Logger.Fatal("创建主题失败 %s: %v" + topic + err.Error())
		}
		global.Logger.Info("Topic %s 成功备份" + topic)
	} else {
		global.Logger.Info("Topic %s 已经存在" + topic)
	}
}

func SetupKafka() {
	// 创建用户主题名称
	topic := "RegisterReq"
	brokers := global.Config.KafkaConfig.Brokers
	global.KafkaBrokers = brokers

	// 创建主题
	CreateTopic(global.KafkaBrokers, topic, 5, 3)

}
