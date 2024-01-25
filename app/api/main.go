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

	//mysql.MakeAccountList(global.MysqlDB)
	//mysql.MakeAllProductList(global.MysqlDB)
	//mysql.MakeHostList(global.MysqlDB)
	//mysql.AddAccount(global.MysqlDB, "蒋", "123", "boss", "Eode")
	mysql.FindUser(global.MysqlDB, "jiang", "AccountList")
	//mysql.AddProduct(global.MysqlDB, "蒋", "可乐", 3, 2.50)
	//mysql.AddProduct(global.MysqlDB, "蒋", "冰可乐", 2, 3.00)
	//mysql.ShowProductList(global.MysqlDB, "兰")
	//mysql.AddAccount(global.MysqlDB, "兰", "123", "customer", "lanfanya")
	//mysql.AddBalance(global.MysqlDB, "兰", 100)

	//mysql.UpdateAllProductList(global.MysqlDB)
	//mysql.ShowAllProductList(global.MysqlDB)
	//mysql.BuyMakeCar(global.MysqlDB)
	//mysql.MakeHostList(global.MysqlDB)
	//mysql.PrintHotList(global.MysqlDB)

	//mysql.BuyAddCar(global.MysqlDB, 3, 1, "兰")
	//mysql.BuyAddCar(global.MysqlDB, 4, 1, "兰")
	//mysql.ShowAllProductList(global.MysqlDB)
	//mysql.ShowCarList(global.MysqlDB)

	//mysql.FindUser(global.MysqlDB, "蒋", "AccountList")
	//mysql.BuyEmptyCar(global.MysqlDB, "兰")
	//mysql.FindUser(global.MysqlDB, "蒋", "AccountList")
	//mysql.FindUser(global.MysqlDB, "兰", "AccountList")
	//mysql.ShowProductList(global.MysqlDB, "兰")
	//mysql.ShowCarList(global.MysqlDB)

	//mysql.DropList(global.MysqlDB, "HostList")
	//mysql.DropList(global.MysqlDB, "jiang")
	//mysql.DropList(global.MysqlDB, "lan")
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
