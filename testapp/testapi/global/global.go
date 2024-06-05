package global

import (
	"database/sql"
	"dianshang/testapp/testapi/global/config"
	"github.com/IBM/sarama"
	"github.com/go-redis/redis/v8"
	"github.com/samuel/go-zookeeper/zk"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"net/http"
)

// 这个文件用来定义一些全局变量
var ( //这个全局变量用来记录日志，将Config和Logger定义为全局变量，以便在整个应用程序中方便地访问它们。
	// Config 这样定义允许您在不同的包和函数中使用它们，确保日志记录器和配置信息对整个应用程序可见。
	//config这个全局变量通常用来存储应用程序的配置信息。在应用程序启动时，会从配置文件中读取配置信息，并将其填充到 Config 结构体中。
	//在整个应用程序中，其他模块或函数可以访问 Config 变量来获取应用程序的配置信息，比如数据库连接信息、日志设置等。
	Config        *config.Config
	Logger        *zap.Logger
	MysqlDB       *sql.DB
	RedisDB       *redis.Client
	EtcdClient    *clientv3.Client
	KafkaProducer sarama.SyncProducer
	ZookeeperConn *zk.Conn
	JaegerClient  *http.Client
	NginxClient   *http.Client
)
