package main

import (
	"dianshang/testapp/testapi/global"
	"dianshang/testapp/testapi/internal/initialize"
	"dianshang/testapp/testapi/router"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"strconv"
)

func main() {
	initialize.SetupViper()
	initialize.SetupLogger()
	initialize.SetupDataBase()
	initialize.SetupEtcd()
	initialize.SetupKafka()
	fmt.Println(global.KafkaProducer)
	initialize.SetupZookeeper()

	initialize.SetupJaeger()
	initialize.SetupNginx()
	initialize.Check()
	config := global.Config.ServerConfig

	// 设置 Gin 模式
	gin.SetMode(config.Mode)
	portStr := strconv.Itoa(config.Port)
	global.Logger.Info("初始化服务器成功", zap.String("port", config.Host+":"+portStr))
	err := router.InitRouter(portStr)
	if err != nil {
		global.Logger.Fatal("服务器启动失败," + err.Error())
	}

}
