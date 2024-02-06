package main

import (
	"dianshang/app/api/global"
	"dianshang/app/api/internal/dao/mysql"
	"dianshang/app/api/internal/initialize"
	"dianshang/app/api/router"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func main() {
	initialize.SetupViper()
	initialize.SetupLogger()
	initialize.SetupDataBase()
	config := global.Config.ServerConfig

	mysql.MakeAccountList(global.MysqlDB)
	mysql.MakeAllProductList(global.MysqlDB)
	mysql.MakeHostList(global.MysqlDB)

	//mysql.DropList(global.MysqlDB, "HostList")
	//mysql.DropList(global.MysqlDB, "jiang")
	//mysql.DropList(global.MysqlDB, "lan")
	//mysql.DropList(global.MysqlDB, "wen")
	//mysql.DropList(global.MysqlDB, "AllProductList")
	//mysql.DropList(global.MysqlDB, "AccountList")
	//mysql.DropList(global.MysqlDB, "Car")
	// 设置 Gin 模式
	gin.SetMode(config.Mode)
	global.Logger.Info("初始化服务器成功", zap.String("port", config.Port+":"+config.Port))
	err := router.InitRouter(config.Port)
	if err != nil {
		global.Logger.Fatal("服务器启动失败," + err.Error())

	}

}
