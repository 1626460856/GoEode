package initialize

import (
	"database/sql"
	"dianshang/app/api/global"
)

// SetupDataBase 这个文件用来描写登录数据库的方法
// 下面是docker中创建mysql数据库的指令
// docker run -d --name=mysql-container -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=test -e TZ=Asia/Shanghai -e MYSQL_CHARSET=utf8mb4 mysql:latest

// docker run -d --name=mysql-container -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=test -e TZ=Asia/Shanghai -e MYSQL_CHARSET=utf8mb4 mysql:latest

func SetupDataBase() {
	SetupMysql()

}

func SetupMysql() {
	config := global.Config.DatabaseConfig.MysqlConfig

	db, err := sql.Open("mysql", config.GetDsn())
	if err != nil {
		global.Logger.Fatal("open mysql failed," + err.Error())
	}
	db.SetConnMaxLifetime(config.ConnMaxLifetime) //最长连接时间
	db.SetConnMaxIdleTime(config.ConnMaxIdleTime) //最长空闲时间
	db.SetMaxIdleConns(config.MaxIdleConns)       //最长控制时间
	db.SetMaxOpenConns(config.MaxOpenConns)
	err = db.Ping()
	if err != nil {
		global.Logger.Fatal("connect to mysql failed ," + err.Error())

	}
	global.MysqlDB = db //赋值给了全局变量
	global.Logger.Info("init mysql success")
}
