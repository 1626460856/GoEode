package svc

import (
	"dianshang/testapp/User/internal/config"
)

// ServiceContext 这个结构体包含应用程序的全局配置和 Kafka 客户端实例
type ServiceContext struct {
	Config      config.Config
	KafkaClient *KafkaClient
}

// NewServiceContext 用于创建并初始化一个新的 ServiceContext 实例
func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		KafkaClient: NewKafkaClient(c.Kafka.Brokers),
	}
}
