package svc

import (
	"github.com/IBM/sarama"
)

type KafkaClient struct {
	Producer sarama.SyncProducer
}

// NewKafkaClient 用于创建并初始化一个新的 KafkaClient 实例
func NewKafkaClient(brokers []string) *KafkaClient {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		//global.Logger.Fatal("启动Kafka producer失败: %v" + err.Error())
	}

	return &KafkaClient{
		Producer: producer,
	}
}

// SendMessage 用于发送消息到指定的 Kafka 主题
func (k *KafkaClient) SendMessage(topic string, message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}
	_, _, err := k.Producer.SendMessage(msg)
	return err
}
func (k *KafkaClient) SendMessage2(topic string, key string, message []byte, headers map[string]string) error {
	var saramaHeaders []sarama.RecordHeader
	for k, v := range headers {
		saramaHeaders = append(saramaHeaders, sarama.RecordHeader{
			Key:   []byte(k),
			Value: []byte(v),
		})
	}

	msg := &sarama.ProducerMessage{
		Topic:   topic,
		Key:     sarama.StringEncoder(key),
		Value:   sarama.ByteEncoder(message),
		Headers: saramaHeaders,
	}

	_, _, err := k.Producer.SendMessage(msg)
	return err
}
