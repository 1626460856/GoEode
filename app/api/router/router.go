package router

import (
	"dianshang/app/api/internal/middleware"
	"dianshang/app/api/internal/service"
	"github.com/gin-gonic/gin"
)

func InitRouter(port string) error {
	r := gin.Default()
	r.Use(middleware.CORS())

	r.POST("/register", service.AddAccountHandler)
	//这是对创建用户api的模拟调用
	// Invoke-WebRequest -Uri http://localhost:8080/register -Method POST -Body "useraccount=jiang&password=123&identity=boss&nickname=Eode"

	// Invoke-WebRequest -Uri http://localhost:8080/register -Method POST -Body "useraccount=lan&password=123&identity=customer&nickname=Lan"

	r.POST("/login", service.Login)
	//模拟登录
	// Invoke-WebRequest -Uri "http://localhost:8080/login" -Method POST -Body "useraccount=jiang&password=123"

	//Invoke-WebRequest -Uri "http://localhost:8080/login" -Method POST -Body "useraccount=lan&password=123"

	r.POST("/findProduct", service.FindProduct)
	//满足了模糊查找，比如可以传入“乐”
	//# 设置商品名称
	//$PRODUCT_NAME = "乐"
	//
	//# 发送 POST 请求，包含商品名称
	//$body = @{
	//   ProductName = $PRODUCT_NAME
	//} | ConvertTo-Json
	//
	//$response = Invoke-RestMethod -Uri "http://localhost:8080/findProduct" -Method Post -Headers @{"Content-Type"="application/json"} -Body $body
	//$response
	r.Use(middleware.JWTAuthMiddleware()) //这里开始以后的调用令牌保护
	r.POST("/call", service.Call)
	// Invoke-WebRequest -Uri "http://localhost:8080/call" -Method POST -Headers @{"Authorization"="Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Ij8iLCJleHAiOjE3MDYwMDkwNzgsImlzcyI6IuiSi-WNk-eHgyJ9.GMi1FUgLHxNo8wqmgp6nfiKxhkQsOdErf-Pl9lEu4b0"}

	r.POST("/myself", service.Myself)
	//$headers = @{
	//	Authorization = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxhbiIsImV4cCI6MTcwNjUxNjQ0MCwiaXNzIjoi6JKL5Y2T54eDIn0.duaM63t6bXKG6HjCrPbQnKH74pJP_tgpvjIrZOx8Acs"
	//}
	//
	//$body = @{
	//	useraccount = "lan"
	//}
	//
	//Invoke-RestMethod -Uri "http://localhost:8080/myself" -Method Post -Headers $headers -Body ($body | ConvertTo-Json)

	r.POST("/refresh", service.Refresh)
	//$webRequest = Invoke-WebRequest -Uri "http://localhost:8080/refresh" -Method POST -Headers @{"Authorization"="Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxhbiIsImV4cCI6MTcwNjUxNjQ0MCwiaXNzIjoi6JKL5Y2T54eDIn0.duaM63t6bXKG6HjCrPbQnKH74pJP_tgpvjIrZOx8Acs"}
	//$webRequest.Content
	r.POST("/addProduct", service.AddProduct)
	//先登录获取令牌
	//$response = Invoke-WebRequest -Uri "http://localhost:8080/login" -Method POST -Body "useraccount=jiang&password=123"
	//$token = ($response.Content | ConvertFrom-Json).token
	//再添加商品
	//$headers = @{"Authorization"="Bearer $token"}
	//$body = @{userAccount="wen"; productName="牛奶"; productNumber=3; productPrice=3.50}
	//Invoke-RestMethod -Uri "http://localhost:8080/addProduct" -Method Post -Headers $headers -Body $body
	r.POST("/addBalance", service.AddBalance)
	//$headers = @{
	//  "Authorization" = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxhbiIsImV4cCI6MTcwNjE5MDI3NywiaXNzIjoi6JKL5Y2T54eDIn0.6RPAz34UpmAASmdyozguzO4HzxhQ0yM9hL5CJTJeIj4"
	//}
	//
	//$body = @{
	//  "useraccount" = "lan"
	//  "money" = 100
	//} | ConvertTo-Json
	//
	//Invoke-WebRequest -Uri "http://localhost:8080/addBalance" -Method POST -Headers $headers -Body $body -ContentType "application/json"
	//
	r.POST("/makeCar", service.MakeCar)
	//$headers = @{
	//   "Authorization" = "Bearer  eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxhbiIsImV4cCI6MTcwNjUxNjQ0MCwiaXNzIjoi6JKL5Y2T54eDIn0.duaM63t6bXKG6HjCrPbQnKH74pJP_tgpvjIrZOx8Acs"
	//}
	//
	//Invoke-WebRequest -Uri "http://localhost:8080/makeCar" -Method POST -Headers $headers -ContentType "application/json"
	r.POST("/addCar", service.AddCar)
	//$headers = @{
	//"Authorization" = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxhbiIsImV4cCI6MTcwNjUxNjQ0MCwiaXNzIjoi6JKL5Y2T54eDIn0.duaM63t6bXKG6HjCrPbQnKH74pJP_tgpvjIrZOx8Acs"
	//}
	//
	//$body = @{
	//"BuyAllProductID" = 5
	//"BuyProductNumber" = 1
	//"BuyUserAccount" = "lan"
	//} | ConvertTo-Json
	//
	//Invoke-WebRequest -Uri "http://localhost:8080/addCar" -Method POST -Headers $headers -Body $body -ContentType "application/json"
	//
	r.POST("/lookCar", service.LookCar)
	//$headers = @{
	// "Authorization" = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxhbiIsImV4cCI6MTcwNjUxNjQ0MCwiaXNzIjoi6JKL5Y2T54eDIn0.duaM63t6bXKG6HjCrPbQnKH74pJP_tgpvjIrZOx8Acs"
	// "Content-Type" = "application/json"
	//}
	//
	//$response = Invoke-WebRequest -Uri "http://localhost:8080/lookCar" -Method POST -Headers $headers
	//$carList = $response.Content | ConvertFrom-Json
	//$carList
	r.POST("/emptyCar", service.EmptyCar)
	//$headers = @{
	//  "Authorization" = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxhbiIsImV4cCI6MTcwNjUxNjQ0MCwiaXNzIjoi6JKL5Y2T54eDIn0.duaM63t6bXKG6HjCrPbQnKH74pJP_tgpvjIrZOx8Acs"
	//  "Content-Type" = "application/json"
	//}
	//
	//$response = Invoke-WebRequest -Uri "http://localhost:8080/emptyCar" -Method POST -Headers $headers
	//$response.Content
	r.POST("/myProductList", service.MyProductList)
	//$headers = @{
	//  "Authorization" = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxhbiIsImV4cCI6MTcwNjUxNjQ0MCwiaXNzIjoi6JKL5Y2T54eDIn0.duaM63t6bXKG6HjCrPbQnKH74pJP_tgpvjIrZOx8Acs"
	//}
	//
	//$body = @{
	//  "useraccount" = "lan"
	//}
	//
	//Invoke-WebRequest -Uri "http://localhost:8080/myProductList" -Method POST -Headers $headers -Body $body
	r.POST("/lookHostList", service.LookHostList)
	//$headers = @{
	//	   "Authorization" = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxhbiIsImV4cCI6MTcwNjUxNjQ0MCwiaXNzIjoi6JKL5Y2T54eDIn0.duaM63t6bXKG6HjCrPbQnKH74pJP_tgpvjIrZOx8Acs"
	//	   "Content-Type" = "application/json"
	//	}
	//
	//	$response = Invoke-WebRequest -Uri "http://localhost:8080/lookHostList" -Method POST -Headers $headers
	//	$HostList = $response.Content | ConvertFrom-Json
	//	$HostList
	r.POST("/changeAccount", service.ChangeAccount)
	//修改的位置可以是"Password"可以是"Nickname"可以是"Identity"其中"Identity"只能是修改为"boss"或者"customer"，另外两个的修改似乎不能输入中文，我不知道为什么
	//$headers = @{
	//  "Authorization" = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxhbiIsImV4cCI6MTcwNjE5MDI3NywiaXNzIjoi6JKL5Y2T54eDIn0.6RPAz34UpmAASmdyozguzO4HzxhQ0yM9hL5CJTJeIj4"
	//}
	//
	//$body = @{
	//  UserAccount = "lan"
	//  Changelocation = "Nickname"
	//  Changetext = "lanfanya"
	//}
	//
	//Invoke-RestMethod -Uri "http://localhost:8080/changeAccount" -Method Post -Headers $headers -Body ($body | ConvertTo-Json)

	err := r.Run(":" + port)
	if err != nil {
		return err
	}
	return err
}
