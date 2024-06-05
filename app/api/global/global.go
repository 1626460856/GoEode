package global

import (
	"database/sql"
	"dianshang/app/api/global/config"
	"github.com/go-redis/redis/v8"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

// 这个文件用来定义一些全局变量
var ( //这个全局变量用来记录日志，将Config和Logger定义为全局变量，以便在整个应用程序中方便地访问它们。
	// Config 这样定义允许您在不同的包和函数中使用它们，确保日志记录器和配置信息对整个应用程序可见。
	//config这个全局变量通常用来存储应用程序的配置信息。在应用程序启动时，会从配置文件中读取配置信息，并将其填充到 Config 结构体中。
	//在整个应用程序中，其他模块或函数可以访问 Config 变量来获取应用程序的配置信息，比如数据库连接信息、日志设置等。
	Config *config.Config
	// Logger 这是一个日志记录器对象，通常用于记录应用程序的日志信息。
	//通过在全局变量中存储日志记录器对象，可以在整个应用程序中方便地使用它来记录各种事件、错误和调试信息。
	Logger *zap.Logger
	// MysqlDB 这个全局变量通常用来存储应用程序与 MySQL 数据库的连接。在应用程序初始化时，会创建一个数据库连接并将其存储在这个全局变量中。
	//这样，在应用程序的其他部分，无需每次都重新创建数据库连接，可以直接使用已经建立好的全局连接对象。
	MysqlDB    *sql.DB
	RedisDB    *redis.Client
	EtcdClient *clientv3.Client
)
