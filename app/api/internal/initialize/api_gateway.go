package initialize

import (
	"context"
	"dianshang/app/api/global"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func SetApiGateway() {
	SetupEtcd()
}
func SetupEtcd() {
	//etcdconfig := global.Config.EtcdConfig
	//dsn := etcdconfig.GetDsn()
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"}, // 更新为 Docker 中运行的地址

		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		global.Logger.Fatal("连接到etcd失败: " + err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = cli.Put(ctx, "test_key", "test_value")
	if err != nil {
		global.Logger.Fatal("向etcd写入数据失败: " + err.Error())
	}

	global.EtcdClient = cli

	// 记录初始化成功消息
	global.Logger.Info("初始化etcd成功")
}
