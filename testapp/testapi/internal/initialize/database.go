package initialize

import (
	"context"
	"database/sql"
	"dianshang/testapp/testapi/global"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

// SetupDataBase 这个文件用来描写登录数据库的方法
// 下面是docker中创建mysql数据库的指令
// docker run -d --name=mysql-container -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=test -e TZ=Asia/Shanghai -e MYSQL_CHARSET=utf8mb4 mysql:latest

// docker run -d --name=Eodemysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=test -e TZ=Asia/Shanghai -e MYSQL_CHARSET=utf8mb4 -v Eode_mysql:/var/lib/mysql mysql:latest

//下面是docker中创建etcd容器的指令
//docker run -d --name=Eodeetcd -p 2379:2379   -v Eode_etcd:/etcd-data bitnami/etcd:latest

//下面是docker中创建redis的指令
//docker run -d --name Eoderedis -p 6379:6379 --restart unless-stopped -v Eode_redis:/data -v /home/redis/conf/redis.conf:/etc/redis/redis.conf redis:latest redis-server /etc/redis/redis.conf --requirepass 123awzsex

func SetupDataBase() {
	SetupMysql()
	SetupRedis()

}

func SetupMysql() {
	config := global.Config.UserdataConfig.MysqlConfig

	db, err := sql.Open("mysql", config.GetDsn())
	if err != nil {
		global.Logger.Fatal("open usermysql failed," + err.Error())
	}
	db.SetConnMaxLifetime(config.ConnMaxLifeTime) //最长连接时间
	db.SetConnMaxIdleTime(config.ConnMaxIdleTime) //最长空闲时间
	db.SetMaxIdleConns(config.MaxIdleConns)       //最长控制时间
	db.SetMaxOpenConns(config.MaxOpenConns)
	err = db.Ping()
	if err != nil {
		global.Logger.Fatal("connect to mysql failed ," + err.Error())

	}
	global.UserMysqlDB = db //赋值给了全局变量
	global.Logger.Info("init usermysql success")

	config = global.Config.ShopdataConfig.MysqlConfig

	db, err = sql.Open("mysql", config.GetDsn())
	if err != nil {
		global.Logger.Fatal("open shopmysql failed," + err.Error())
	}
	db.SetConnMaxLifetime(config.ConnMaxLifeTime) //最长连接时间
	db.SetConnMaxIdleTime(config.ConnMaxIdleTime) //最长空闲时间
	db.SetMaxIdleConns(config.MaxIdleConns)       //最长控制时间
	db.SetMaxOpenConns(config.MaxOpenConns)
	err = db.Ping()
	if err != nil {
		global.Logger.Fatal("connect to shopmysql failed ," + err.Error())

	}
	global.ShopMysqlDB = db //赋值给了全局变量
	global.Logger.Info("init shopmysql success")
}
func SetupRedis() {
	// 从全局配置中获取Redis数据库的配置信息
	config := global.Config.UserdataConfig.Redis1Config

	// 创建一个Redis客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.GetDsn(), // 使用 GetDsn() 方法生成连接字符串
		Password: config.Password, // 设置Redis密码
	})

	// 使用Ping操作检查与Redis1服务器的连接
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Fatal("连接到userRedis1失败: " + err.Error())
	}

	// 将Redis客户端赋值给全局变量
	global.UserRedis1DB = rdb

	// 记录初始化成功消息
	global.Logger.Info("初始化userRedis1成功")
	// 从全局配置中获取Redis数据库的配置信息
	config2 := global.Config.UserdataConfig.Redis2Config

	// 创建一个Redis客户端
	rdb = redis.NewClient(&redis.Options{
		Addr:     config2.GetDsn(), // 使用 GetDsn() 方法生成连接字符串
		Password: config2.Password, // 设置Redis密码
	})

	// 使用Ping操作检查与Redis2服务器的连接
	ctx = context.Background()
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Fatal("连接到userRedis2失败: " + err.Error())
	}

	// 将Redis客户端赋值给全局变量
	global.UserRedis2DB = rdb

	// 记录初始化成功消息
	global.Logger.Info("初始化userRedis2成功")

	// 从全局配置中获取Redis1数据库的配置信息
	config = global.Config.ShopdataConfig.Redis1Config

	// 创建一个Redis客户端
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.GetDsn(), // 使用 GetDsn() 方法生成连接字符串
		Password: config.Password, // 设置Redis密码
	})

	// 使用Ping操作检查与Redis服务器的连接
	ctx = context.Background()
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Fatal("连接到shopRedis1失败: " + err.Error())
	}

	// 将Redis客户端赋值给全局变量
	global.ShopRedis1DB = rdb

	// 记录初始化成功消息
	global.Logger.Info("初始化shopRedis1成功")
	// 从全局配置中获取Redis2数据库的配置信息
	config2 = global.Config.ShopdataConfig.Redis2Config

	// 创建一个Redis客户端
	rdb = redis.NewClient(&redis.Options{
		Addr:     config2.GetDsn(), // 使用 GetDsn() 方法生成连接字符串
		Password: config2.Password, // 设置Redis密码
	})

	// 使用Ping操作检查与Redis服务器的连接
	ctx = context.Background()
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Fatal("连接到shopRedis2失败: " + err.Error())
	}

	// 将Redis客户端赋值给全局变量
	global.ShopRedis2DB = rdb

	// 记录初始化成功消息
	global.Logger.Info("初始化shopRedis2成功")
}
