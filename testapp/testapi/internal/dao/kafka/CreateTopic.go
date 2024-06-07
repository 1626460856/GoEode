package kafka

import (
	"dianshang/testapp/testapi/global"
	"github.com/IBM/sarama"
)

// CreateTopic 创建主题函数传入参数：
// 1.kafka节点
// 2.主题的名称
// 2.主题内部的分区数量
// 3.主题保存的份数，如果是3则三个节点都保存数据
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
