package initialize

import (
	"dianshang/testapp/testapi/global"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

func SetupZookeeper() {
	// 从全局配置中获取Zookeeper的配置信息
	config := global.Config.ZookeeperConfig

	// 创建一个Zookeeper客户端
	conn, _, err := zk.Connect([]string{config.Address}, time.Second*10)
	if err != nil {
		global.Logger.Fatal("连接到Zookeeper失败: " + err.Error())
	}

	// 检查与Zookeeper服务器的连接状态
	_, _, err = conn.Get("/brokers/ids")
	if err != nil {
		global.Logger.Fatal("Zookeeper健康检查失败: " + err.Error())
	}

	// 将Zookeeper连接赋值给全局变量
	global.ZookeeperConn = conn

	// 记录初始化成功消息
	global.Logger.Info("初始化Zookeeper成功")
}
