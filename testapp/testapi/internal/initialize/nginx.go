package initialize

import (
	"dianshang/testapp/testapi/global"
	"fmt"
	"net/http"
	"time"
)

func SetupNginx() {
	// 从全局配置中获取Nginx的配置信息
	config := global.Config.NginxConfig

	// 创建一个Nginx客户端
	nginxClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	// 测试与Nginx服务器的连接
	resp, err := nginxClient.Get(fmt.Sprintf("http://%s:%d", config.Address, config.Port))
	if err != nil || resp.StatusCode != http.StatusOK {
		global.Logger.Fatal("连接到Nginx失败: " + err.Error())
	}
	defer resp.Body.Close()

	// 将Nginx客户端赋值给全局变量
	global.NginxClient = nginxClient

	// 记录初始化成功消息
	global.Logger.Info("初始化Nginx成功")
}
