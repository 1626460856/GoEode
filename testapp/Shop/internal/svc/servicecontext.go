package svc

import (
	"dianshang/testapp/Shop/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	KafkaClient *KafkaClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		KafkaClient: NewKafkaClient(c.Kafka.Brokers),
	}
}
