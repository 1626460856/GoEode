package initialize

import (
	"context"
	"dianshang/testapp/testapi/global"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"log"
	"net/http"
)

func Check() {
	// 检查 Config
	if global.Config == nil {
		log.Fatal("Config 未初始化")
	} else {
		fmt.Println("Config 已初始化")
	}

	// 检查 Logger
	if global.Logger == nil {
		log.Fatal("Logger 未初始化")
	} else {
		global.Logger.Info("Logger 已初始化")
	}

	// 检查 MySQL 连接
	if global.MysqlDB == nil {
		log.Fatal("MysqlDB 未初始化")
	} else {
		err := global.MysqlDB.Ping()
		if err != nil {
			log.Fatal("无法连接到 MySQL:", err)
		} else {
			fmt.Println("MysqlDB 连接成功")
		}
	}

	// 检查 Redis 连接
	if global.RedisDB == nil {
		log.Fatal("RedisDB 未初始化")
	} else {
		ctx := context.Background()
		_, err := global.RedisDB.Ping(ctx).Result()
		if err != nil {
			log.Fatal("无法连接到 Redis:", err)
		} else {
			fmt.Println("RedisDB 连接成功")
		}
	}

	// 检查 Etcd 连接
	if global.EtcdClient == nil {
		log.Fatal("EtcdClient 未初始化")
	} else {
		ctx := context.Background()
		_, err := global.EtcdClient.Get(ctx, "test-key")
		if err != nil {
			log.Fatal("无法连接到 Etcd:", err)
		} else {
			fmt.Println("EtcdClient 连接成功")
		}
	}
	// 检查 Nginx 客户端
	if global.NginxClient == nil {
		log.Fatal("NginxClient 未初始化")
	} else {
		resp, err := global.NginxClient.Get(fmt.Sprintf("http://%s:%d", global.Config.NginxConfig.Address, global.Config.NginxConfig.Port))
		if err != nil || resp.StatusCode != http.StatusOK {
			log.Fatal("无法连接到 Nginx:", err)
		} else {
			defer resp.Body.Close()
			fmt.Println("NginxClient 连接成功")
		}
	}
	// 检查 Zookeeper 连接
	conn := global.ZookeeperConn
	state := conn.State()
	if state == zk.StateHasSession {
		global.Logger.Info("Zookeeper 连接正常")
	} else {
		global.Logger.Fatal("Zookeeper 连接检查失败，状态: " + state.String())
	}
	// 检查 Jaeger 客户端
	if global.JaegerClient == nil {
		log.Fatal("JaegerClient 未初始化")
	} else {
		resp, err := global.JaegerClient.Get(global.Config.JaegerConfig.Endpoint)
		if err != nil || resp.StatusCode != http.StatusOK {
			log.Fatal("无法连接到 Jaeger:", err)
		} else {
			defer resp.Body.Close()
			fmt.Println("JaegerClient 连接成功")
		}
	}
	// 检查 Kafka 生产者
	fmt.Println(global.KafkaProducer)
}
