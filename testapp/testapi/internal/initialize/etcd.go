package initialize

import (
	"context"
	"dianshang/testapp/testapi/global"
	"go.etcd.io/etcd/client/v3"
	"time"
)

func SetupEtcd() {
	// 从全局配置中获取Etcd的配置信息
	config := global.Config.EtcdConfig

	// 创建一个Etcd客户端
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   config.Endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		global.Logger.Fatal("连接到Etcd失败: " + err.Error())
	}

	// 使用健康检查来验证与Etcd服务器的连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = etcdClient.Status(ctx, config.Endpoints[0])
	if err != nil {
		global.Logger.Fatal("Etcd健康检查失败: " + err.Error())
	}

	// 将Etcd客户端赋值给全局变量
	global.EtcdClient = etcdClient

	// 记录初始化成功消息
	global.Logger.Info("初始化Etcd成功")
}
