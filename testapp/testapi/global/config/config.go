package config

import (
	"fmt"
	"time"
)

// Config 这个文件用来描述配置文件的结构体
// 在代码中，标签 mapstructure:"zap_config" json:"zap_config" 和 mapstructure:"database" json:"database"
// 表示 ZapConfig 和 DatabaseConfig 结构体分别对应配置文件中的 zap_config 和 database 部分。

// Config 这个文件用来描述配置文件的结构体
type Config struct {
	ZapConfig       ZapConfig       `mapstructure:"zap_config" json:"zap_config"`
	DatabaseConfig  DatabaseConfig  `mapstructure:"database" json:"database"`
	ServerConfig    ServerConfig    `mapstructure:"server" json:"server"`
	EtcdConfig      EtcdConfig      `mapstructure:"etcd" json:"etcd"`
	ZookeeperConfig ZookeeperConfig `mapstructure:"zookeeper" json:"zookeeper"`
	KafkaConfig     KafkaConfig     `mapstructure:"kafka" json:"kafka"`
	JaegerConfig    JaegerConfig    `mapstructure:"jaeger" json:"jaeger"`
	NginxConfig     NginxConfig     `mapstructure:"nginx" json:"nginx"`
}

// DatabaseConfig 结构体用来描述数据库配置
type DatabaseConfig struct {
	MysqlConfig `mapstructure:"mysql" json:"mysql"`
	RedisConfig `mapstructure:"redis" json:"redis"`
}

// ServerConfig 结构体用来描述服务器配置
type ServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Mode string `mapstructure:"mode" json:"mode"`
}

// ZapConfig 结构体用来描述日志配置
type ZapConfig struct {
	Filename   string `mapstructure:"filename" json:"filename"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size"`
	MaxAge     int    `mapstructure:"max_age" json:"max_age"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups"`
}

type ZookeeperConfig struct {
	Address string `mapstructure:"address"`
}

type KafkaConfig struct {
	Brokers []string `mapstructure:"brokers"`
}

type JaegerConfig struct {
	Endpoint string `mapstructure:"endpoint"`
}

type NginxConfig struct {
	Address string `mapstructure:"address"`
	Port    int    `mapstructure:"port"`
}

// MysqlConfig 结构体用来描述MySQL数据库配置
type MysqlConfig struct {
	Addr            string        `mapstructure:"addr"`
	Port            int           `mapstructure:"port"`
	Username        string        `mapstructure:"username"`
	Password        string        `mapstructure:"password"`
	DB              string        `mapstructure:"db"`
	Charset         string        `mapstructure:"charset"`
	ConnMaxIdleTime time.Duration `mapstructure:"connMaxIdleTime"`
	ConnMaxLifeTime time.Duration `mapstructure:"connMaxLifeTime"`
	Place           string        `mapstructure:"place"`
	MaxIdleConns    int           `mapstructure:"maxIdleConns"`
	MaxOpenConns    int           `mapstructure:"maxOpenConns"`
}

// RedisConfig 结构体用来描述Redis数据库配置
type RedisConfig struct {
	Address         string        `mapstructure:"address"`
	Port            int           `mapstructure:"port"`
	Password        string        `mapstructure:"password"`
	ConnMaxIdleTime time.Duration `mapstructure:"connMaxIdleTime"`
	ConnMaxLifeTime time.Duration `mapstructure:"connMaxLifeTime"`
	MaxIdleConns    int           `mapstructure:"maxIdleConns"`
	MaxOpenConns    int           `mapstructure:"maxOpenConns"`
}

type EtcdConfig struct {
	Endpoints []string `mapstructure:"endpoints"`
}

func (r *RedisConfig) GetDsn() string {
	return fmt.Sprintf("%s:%d", r.Address, r.Port)
}

func (m *MysqlConfig) GetDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s",
		m.Username,
		m.Password,
		m.Addr,
		m.Port,
		m.DB,
		m.Charset,
		m.Place,
	)
}
