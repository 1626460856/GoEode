package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string `json:"access_secret"`
		AccessExpire int64  `json:"access_expire"`
	}
	Kafka struct {
		Brokers []string
	}
}
