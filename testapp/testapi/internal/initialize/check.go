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

	// 检查 UserMySQL 连接
	if global.UserMysqlDB == nil {
		log.Fatal("UserMysqlDB 未初始化")
	} else {
		err := global.UserMysqlDB.Ping()
		if err != nil {
			log.Fatal("无法连接到 UserMySQL:", err)
		} else {
			fmt.Println("UserMysqlDB 连接成功")
		}
	}
	// 检查 ShopMySQL 连接
	if global.ShopMysqlDB == nil {
		log.Fatal("ShopMysqlDB 未初始化")
	} else {
		err := global.ShopMysqlDB.Ping()
		if err != nil {
			log.Fatal("无法连接到 ShopMySQL:", err)
		} else {
			fmt.Println("ShopMysqlDB 连接成功")
		}
	}

	// 检查 UserRedis1 连接
	if global.UserRedis1DB == nil {
		log.Fatal("UserRedis1DB 未初始化")
	} else {
		ctx := context.Background()
		_, err := global.UserRedis1DB.Ping(ctx).Result()
		if err != nil {
			log.Fatal("无法连接到 UserRedis1:", err)
		} else {
			fmt.Println("UserRedis1DB 连接成功")
		}
	}
	// 检查 ShopRedis1 连接
	if global.ShopRedis1DB == nil {
		log.Fatal("ShopRedis1DB 未初始化")
	} else {
		ctx := context.Background()
		_, err := global.ShopRedis1DB.Ping(ctx).Result()
		if err != nil {
			log.Fatal("无法连接到 ShopRedis1:", err)
		} else {
			fmt.Println("ShopRedis1DB 连接成功")
		}
	}
	// 检查 UserRedis2 连接
	if global.UserRedis2DB == nil {
		log.Fatal("UserRedis2DB 未初始化")
	} else {
		ctx := context.Background()
		_, err := global.UserRedis2DB.Ping(ctx).Result()
		if err != nil {
			log.Fatal("无法连接到 UserRedi2:", err)
		} else {
			fmt.Println("UserRedis2DB 连接成功")
		}
	}
	// 检查 ShopRedis2 连接
	if global.ShopRedis2DB == nil {
		log.Fatal("ShopRedis2DB 未初始化")
	} else {
		ctx := context.Background()
		_, err := global.ShopRedis2DB.Ping(ctx).Result()
		if err != nil {
			log.Fatal("无法连接到 ShopRedis2:", err)
		} else {
			fmt.Println("ShopRedis2DB 连接成功")
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
	//// 检查 Nginx 客户端
	//if global.NginxClient == nil {
	//	log.Fatal("NginxClient 未初始化")
	//} else {
	//	resp, err := global.NginxClient.Get(fmt.Sprintf("http://%s:%d", global.Config.NginxConfig.Address, global.Config.NginxConfig.Port))
	//	if err != nil || resp.StatusCode != http.StatusOK {
	//		log.Fatal("无法连接到 Nginx:", err)
	//	} else {
	//		defer resp.Body.Close()
	//		fmt.Println("NginxClient 连接成功")
	//	}
	//}
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

}
