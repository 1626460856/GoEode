package config

import (
	"fmt"
	"time"
)

// Config 这个文件用来描述配置文件的结构体
// 在代码中，标签 mapstructure:"zap_config" json:"zap_config" 和 mapstructure:"database" json:"database"
// 表示 ZapConfig 和 DatabaseConfig 结构体分别对应配置文件中的 zap_config 和 database 部分。

// Config 结构体用来描述配置文件的整体结构
type Config struct { //config单词是配置的意思
	ZapConfig      ZapConfig      `mapstructure:"zap_config" json:"zap_config"`
	DatabaseConfig DatabaseConfig `mapstructure:"database" json:"database"`
	ServerConfig   ServerConfig   `mapstructure:"server" json:"server"`
	EtcdConfig     EtcdConfig     `mapstructure:"etcd_config" json:"etcd-config"`
}

// ServerConfig 结构体用来描述服务器配置
type ServerConfig struct {
	Host string `mapstructure:"addr" json:"addr"`
	Port string `mapstructure:"port" json:"port"`
	Mode string `mapstructure:"mode" json:"mode"`
}

// ZapConfig 结构体用来描述日志配置
type ZapConfig struct {
	Filename   string `mapstructure:"filename" json:"filename"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size"`
	MaxAge     int    `mapstructure:"max_age" json:"max_age"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups"`
}

// DatabaseConfig 结构体用来描述数据库配置
type DatabaseConfig struct {
	MysqlConfig `mapstructure:"mysql" json:"mysql"`
	RedisConfig `mapstructure:"redis" json:"redis"`
}

// MysqlConfig 结构体用来描述MySQL数据库配置
type MysqlConfig struct {
	Addr            string        `mapstructure:"addr" json:"addr"`
	Port            string        `mapstructure:"port" json:"port"`
	DB              string        `mapstructure:"db" json:"db"`
	Username        string        `mapstructure:"username" json:"username"`
	Password        string        `mapstructure:"password" json:"password"`
	ConnMaxLifetime time.Duration `mapstructure:"connMaxLifeTime" json:"connMaxLifeTime"`
	ConnMaxIdleTime time.Duration `mapstructure:"connMaxIdleTime" json:"connMaxIdleTime"`
	MaxIdleConns    int           `mapstructure:"maxIdleConns" json:"maxIdleConns"`
	MaxOpenConns    int           `mapstructure:" maxOpenConns" json:" maxOpenConns"`
	Charset         string        `mapstructure:"charset" json:"charset"`
	Place           string        `mapstructure:"place" json:"place"`
}

// RedisConfig 结构体用来描述Redis数据库配置
type RedisConfig struct {
	Address         string        `mapstructure:"address" json:"address"`
	Password        string        `mapstructure:"password" json:"password"`
	Port            string        `mapstructure:"port" json:"port"`
	ConnMaxIdleTime time.Duration `mapstructure:"connMaxIdleTime" json:"connMaxIdleTime"`
	ConnMaxLifeTime time.Duration `mapstructure:"connMaxLifeTime" json:"connMaxLifeTime"`
	MaxIdleConns    int           `mapstructure:"maxIdleConns" json:"maxIdleConns"`
	MaxOpenConns    int           `mapstructure:"maxOpenConns" json:"maxOpenConns"`
}

type EtcdConfig struct {
	Address         string        `mapstructure:"address" json:"address"`
	Username        string        `mapstructure:"username" json:"username"`
	Password        string        `mapstructure:"password" json:"password"`
	Port            string        `mapstructure:"port" json:"port"`
	ConnMaxIdleTime time.Duration `mapstructure:"connMaxIdleTime" json:"connMaxIdleTime"`
	ConnMaxLifeTime time.Duration `mapstructure:"connMaxLifeTime" json:"connMaxLifeTime"`
	MaxIdleConns    int           `mapstructure:"maxIdleConns" json:"maxIdleConns"`
	MaxOpenConns    int           `mapstructure:"maxOpenConns" json:"maxOpenConns"`
}

func (r *RedisConfig) GetDsn() string {
	return fmt.Sprintf("%s:%s", r.Address, r.Port)
}

func (m *MysqlConfig) GetDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		m.Username,
		m.Password,
		m.Addr,
		m.Port,
		m.DB,
		m.Charset,
		m.Place,
	)
}
