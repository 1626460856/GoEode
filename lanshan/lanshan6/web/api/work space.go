package api

import (
	"GoEode/lanshan/lanshan6/web/api/middleware"
	"GoEode/lanshan/lanshan6/web/dao"
	"GoEode/lanshan/lanshan6/web/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

// 注册函数
func register(c *gin.Context) {

	from := model.User{}
	if err := c.ShouldBind(&from); err != nil {
		fmt.Println(err)
		dao.RespFail(c, "输入请求值绑定错误")
		return
	} //从c的Post请求中读取表单数据并绑定到from上，如果出现错误，则返回错误信息
	//传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")
	//验证用户名是否重复
	flag := dao.FindUser(username)
	fmt.Println(flag)
	if flag { //如果找到了，那么就已经占用，返回错误
		//以JSON格式返回信息
		dao.RespFail(c, "该用户名已经被占用")
	}
	if !flag {
		dao.AddUser(username, password)
		// 以 JSON 格式返回信息
		dao.RespSuccess(c, "add user successful")
	}
	dao.Writedata(dao.Database)

}
func newpassword(data map[string]string, username string, newpassword string) {
	delete(data, username)
	data[username] = newpassword
}
func changepassword(c *gin.Context) {
	from := model.User{}
	if err := c.ShouldBind(&from); err != nil {
		fmt.Println(err)
		dao.RespFail(c, "输入请求值绑定错误")
		return
	} //从c的Post请求中读取表单数据并绑定到from上，如果出现错误，则返回错误信息
	//传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag := dao.FindUser(username)
	// 不存在则退出
	if !flag {
		// 以 JSON 格式返回信息
		dao.RespFail(c, "该用户名不存在")
		return
	}
	if flag {
		if dao.FindPasswordFromUsername(username) == password {
			dao.RespFail(c, "你输入的新密码与你原密码一致")
		}
		if dao.FindPasswordFromUsername(username) != password {
			newpassword(dao.Database, username, password)
			dao.RespSuccess(c, "修改成功")
		}
	}
	dao.Writedata(dao.Database)

}

// 登录界面
func login(c *gin.Context) {

	if err := c.ShouldBind(&model.User{}); err != nil {
		dao.RespFail(c, "输入请求值绑定错误")
		return
	}
	// 传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")
	//验证用户名是否存在
	flag := dao.FindUser(username)
	// 不存在则退出
	if !flag {
		// 以 JSON 格式返回信息
		dao.RespFail(c, "该用户名不存在")
		return
	}
	// 查找正确的密码
	truePassword := dao.FindPasswordFromUsername(username)
	if flag {
		if truePassword != password {
			// 以 JSON 格式返回信息
			dao.RespFail(c, "用户名存在，但是输入密码错误")
			return
		}
		if truePassword == password {
			// 正确则登录成功
			// 创建一个我们自己的声明
			claim := model.MyClaims{
				Username: username, // 自定义字段
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // 过期时间
					Issuer:    "Yxh",                                // 签发人
				},
			}
			// 使用指定的签名方法创建签名对象
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
			// 使用指定的secret签名并获得完整的编码后的字符串token
			tokenString, _ := token.SignedString(middleware.Secret)
			dao.RespSuccess(c, tokenString)
		}
	}
	dao.Writedata(dao.Database)
}

// 从令牌中获取用户名
func getUsernameFromToken(c *gin.Context) {
	username, _ := c.Get("username")
	dao.RespSuccess(c, username.(string))
}

// 初始化路由器
func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS()) //使用中间体
	r.POST("/register", register)
	r.POST("/login", login)
	r.POST("/changepassword", changepassword)
	UserRouter := r.Group("/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.GET("/get", getUsernameFromToken)
	}
	r.Run(":8088")

}
