package main

import "github.com/gin-gonic/gin"

func main() {
	//创建Gin的路由引擎
	r := gin.Default()
	//定义路由和处理函数
	r.GET("/hello", func(c *gin.Context) {
		//c是一个gin.Context对象，它包含了HTTP请求的上下文信息和用于生成HTTP响应的方法。
		c.JSON(200, gin.H{ //c.JSON是gin.Context对象提供的方法，用于返回一个JSON格式的HTTP响应。
			// 200是状态码，表示访问成功，gin.H是一个Gin框架提供的快捷方式，用于创建一个map[string]interface{}类型的数据结构，
			//也就是一个键值对的集合，其中键是字符串，值可以是任意类型。
			"message": "hello,Gin!", //message是名，hello，Gin！是值
		})

	})
	//启动http服务器并监听指定端口，默认在0.0.0.0:8080启动服务
	r.Run(":8080")
	//在本地运行该程序后，可以在浏览器中访问"http://localhost:8080/hello"来获取响应。
}
