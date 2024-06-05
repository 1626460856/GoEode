package initialize

import (
	"dianshang/testapp/testapi/global"
	"net/http"
	"time"
)

func SetupJaeger() {
	// 创建一个Jaeger客户端
	jaegerClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	// 使用配置文件中的 Endpoint 进行健康检查
	resp, err := jaegerClient.Get(global.Config.JaegerConfig.Endpoint)
	if err != nil {
		global.Logger.Fatal("连接到Jaeger失败: " + err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		global.Logger.Fatal("连接到Jaeger失败，状态码: " + resp.Status)
		return
	}

	// 将Jaeger客户端赋值给全局变量
	global.JaegerClient = jaegerClient

	// 记录初始化成功消息
	global.Logger.Info("初始化Jaeger成功")
}
