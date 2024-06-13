package main

import (
	"dianshang/testapp/Pay/database"
	"flag"
	"fmt"

	"dianshang/testapp/Pay/internal/config"
	"dianshang/testapp/Pay/internal/handler"
	"dianshang/testapp/Pay/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "D:/电商/testapp/Pay/etc/shop.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	// 检查UserMySQLDB是否初始化
	err := database.ShopMySQLDB.Ping()
	if err != nil {
		fmt.Printf("Failed to connect to MySQL database: %v\n", err)
		return
	}
	if err == nil {
		fmt.Println("连接数据库成功")
	}
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
