package service

import (
	"dianshang/app/api/global"
	"dianshang/app/api/internal/consts"
	"dianshang/app/api/internal/dao/mysql"
	"dianshang/app/api/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddAccountHandler(c *gin.Context) {
	// 从请求中获取参数
	useraccount := c.PostForm("useraccount")
	password := c.PostForm("password")
	identity := c.PostForm("identity")
	nickname := c.PostForm("nickname")

	// 调用 AddAccount 函数
	err := mysql.AddAccount(global.MysqlDB, useraccount, password, identity, nickname)
	if err != nil {
		global.Logger.Error("创建用户失败，" + err.Error())
		c.JSON(400, gin.H{
			"code": consts.AddAccountFailed, //自定义的错误码，在const's/error_id.go里面
			"msg":  "创建用户失败，" + err.Error(),
		})
		return
	}
	// 处理结果并返回响应
	c.JSON(http.StatusOK, gin.H{"message": "成功注册账户"})
	return
}
func Login(c *gin.Context) {
	userAccount := c.PostForm("useraccount")
	password := c.PostForm("password")

	tokenString, err := mysql.Login(global.MysqlDB, userAccount, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
func Call(c *gin.Context) {
	// 从请求上下文中获取用户名信息
	_, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未经授权的访问"})
		c.Abort()
		return
	}

	// 成功通过了验证，进行其他逻辑处理
	c.JSON(http.StatusOK, gin.H{"message": "成功通过了验证"})
}
func Refresh(c *gin.Context) {
	err := mysql.UpdateAllProductList(global.MysqlDB)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	var allproductlist []model.AllProductList
	allproductlist, _ = mysql.CoutAllProductList(global.MysqlDB)
	c.JSON(http.StatusOK, gin.H{
		"allproductlist": allproductlist,
	})
}
func Myself(c *gin.Context) {
	// 从请求上下文中获取用户名信息
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未经授权的访问"})
		c.Abort()
		return
	}
	userAccount := username.(string) // 假设用户名是字符串类型

	// 成功通过了验证，进行其他逻辑处理
	account, err := mysql.CoutFindUser(global.MysqlDB, userAccount, "AccountList")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取用户信息"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Account": account})
}

func AddProduct(c *gin.Context) {
	// 从请求上下文中获取用户名信息
	_, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未经授权的访问"})
		c.Abort()
		return
	}

	// 成功通过了验证，进行其他逻辑处理
	userAccount := c.PostForm("userAccount")
	productName := c.PostForm("productName")
	productNumberStr := c.DefaultPostForm("productNumber", "0") // 设置默认值为 "0"
	productPriceStr := c.DefaultPostForm("productPrice", "0.0") // 设置默认值为 "0.0"

	// 将字符串转换为目标类型
	productNumber, err := strconv.Atoi(productNumberStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "商品数量必须是整数",
		})
		return
	}

	productPrice, err := strconv.ParseFloat(productPriceStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "商品价格必须是数字",
		})
		return
	}

	err = mysql.AddProduct(global.MysqlDB, userAccount, productName, productNumber, productPrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "成功添加商品到：" + userAccount,
	})
}
func AddBalance(c *gin.Context) {
	// 从请求上下文中获取用户名信息
	_, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未经授权的访问"})
		c.Abort()
		return
	}

	// 从请求中获取充值金额等信息
	var req struct {
		UserAccount string  `json:"useraccount"`
		Money       float64 `json:"money"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	// 调用处理充值逻辑的函数
	err := mysql.AddBalance(global.MysqlDB, req.UserAccount, req.Money)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新账户余额失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "账户余额更新成功"})
}
func MakeCar(c *gin.Context) {
	// 从请求上下文中获取用户名信息
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未经授权的访问"})
		c.Abort()
		return
	}
	//BuyUserAccount := c.PostForm("BuyUserAccount")
	// 调用创建购物车逻辑的函数
	err := mysql.BuyMakeCar(global.MysqlDB, username.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建购物车失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "成功创建购物车"})
}
func AddCar(c *gin.Context) {
	// 从请求上下文中获取用户名信息
	_, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未经授权的访问"})
		c.Abort()
		return
	}
	var buy struct {
		BuyAllProductID  int    `json:"BuyAllProductID"`
		BuyProductNumber int    `json:"BuyProductNumber"`
		BuyUserAccount   string `json:"BuyUserAccount"`
	}
	// 将请求的 JSON 数据绑定到 buy 结构体
	if err := c.ShouldBindJSON(&buy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}

	err := mysql.BuyAddCar(global.MysqlDB, buy.BuyAllProductID, buy.BuyProductNumber, buy.BuyUserAccount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "添加购物车失败" + err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": "成功添加到购物车"})
}
func LookCar(c *gin.Context) {
	// 从请求上下文中获取用户名信息
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未经授权的访问"})
		c.Abort()
		return
	}
	carList, err := mysql.CountCarList(global.MysqlDB, username.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取购物车列表"})
		return
	}

	c.JSON(http.StatusOK, carList)

}
func EmptyCar(c *gin.Context) {
	// 从请求上下文中获取用户名信息
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未经授权的访问"})
		c.Abort()
		return
	}
	err := mysql.BuyEmptyCar(global.MysqlDB, username.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("处理购物车失败: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "成功处理购物车中的商品，并清空了购物车"})
}
func MyProductList(c *gin.Context) {
	// 从请求上下文中获取用户名信息
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未经授权的访问"})
		c.Abort()
		return
	}
	productList, err := mysql.CoutProductList(global.MysqlDB, username.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "无法获取购物车列表" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, productList)
}
func LookHostList(c *gin.Context) {
	// 从请求上下文中获取用户名信息
	_, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未经授权的访问"})
		c.Abort()
		return
	}
	hostlist, err := mysql.SortHotList(global.MysqlDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "无法获取热榜" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"HostList": hostlist,
	})

}
func ChangeAccount(c *gin.Context) {
	// 从请求上下文中获取用户名信息
	_, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未经授权的访问"})
		return
	}

	type Change struct {
		UserAccount    string `json:"UserAccount" `
		Changelocation string `json:"Changelocation" `
		Changetext     string `json:"Changetext" `
	}

	var change Change
	// 将请求的 JSON 数据绑定到 change 结构体
	if err := c.ShouldBindJSON(&change); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}

	err := mysql.ChangeAccount(global.MysqlDB, change.UserAccount, change.Changelocation, change.Changetext)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数错误" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "成功将" + change.UserAccount + "的" + change.Changelocation + "修改为：" + change.Changetext,
	})
}

func FindProduct(c *gin.Context) {
	productName := c.PostForm("ProductName")
	ProductList, err := mysql.CountFindProduct(global.MysqlDB, productName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "获取检索的商品失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ProductList": ProductList,
	})
}
