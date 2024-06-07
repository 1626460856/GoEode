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
	config := global.Config.UserdataConfig.RedisConfig

	// 创建一个Redis客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.GetDsn(), // 使用 GetDsn() 方法生成连接字符串
		Password: config.Password, // 设置Redis密码
	})

	// 使用Ping操作检查与Redis服务器的连接
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Fatal("连接到userRedis失败: " + err.Error())
	}

	// 将Redis客户端赋值给全局变量
	global.UserRedisDB = rdb

	// 记录初始化成功消息
	global.Logger.Info("初始化userRedis成功")

	// 从全局配置中获取Redis数据库的配置信息
	config = global.Config.ShopdataConfig.RedisConfig

	// 创建一个Redis客户端
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.GetDsn(), // 使用 GetDsn() 方法生成连接字符串
		Password: config.Password, // 设置Redis密码
	})

	// 使用Ping操作检查与Redis服务器的连接
	ctx = context.Background()
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Fatal("连接到shopRedis失败: " + err.Error())
	}

	// 将Redis客户端赋值给全局变量
	global.ShopRedisDB = rdb

	// 记录初始化成功消息
	global.Logger.Info("初始化shopRedis成功")
}
