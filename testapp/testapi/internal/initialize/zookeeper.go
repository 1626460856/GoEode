package initialize

import (
	"dianshang/testapp/testapi/global"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

func SetupZookeeper() {
	config := global.Config.ZookeeperConfig
	servers := []string{config.Servers[0] + ":2181"} // 确保格式正确
	fmt.Println(servers)
	conn, _, err := zk.Connect(servers, time.Duration(config.Timeout)*time.Second)
	if err != nil {
		global.Logger.Fatal("连接 Zookeeper 失败: " + err.Error())
	} else {
		global.ZookeeperConn = conn
		global.Logger.Info("连接 Zookeeper 成功")
	}
}
