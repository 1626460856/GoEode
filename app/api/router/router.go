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

	r.Use(middleware.JWTAuthMiddleware()) //这里开始以后的调用令牌保护
	r.POST("/call", service.Call)
	// Invoke-WebRequest -Uri "http://localhost:8080/call" -Method POST -Headers @{"Authorization"="Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Ij8iLCJleHAiOjE3MDYwMDkwNzgsImlzcyI6IuiSi-WNk-eHgyJ9.GMi1FUgLHxNo8wqmgp6nfiKxhkQsOdErf-Pl9lEu4b0"}

	r.POST("/myself", service.Myself)
	//$headers = @{
	//	Authorization = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImppYW5nIiwiZXhwIjoxNzA2MTI4MTc0LCJpc3MiOiLokovljZPnh4MifQ.raZ3eS5jsDGU4yaoWdWCriLR8zMlrsd0fjs5d5rwRcE"
	//}
	//
	//$body = @{
	//	useraccount = "jiang"
	//}
	//
	//Invoke-RestMethod -Uri "http://localhost:8080/myself" -Method Post -Headers $headers -Body ($body | ConvertTo-Json)

	r.POST("/refresh", service.Refresh)
	//$webRequest = Invoke-WebRequest -Uri "http://localhost:8080/refresh" -Method POST -Headers @{"Authorization"="Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxhbiIsImV4cCI6MTcwNjE2NzYzMCwiaXNzIjoi6JKL5Y2T54eDIn0.aWCxS4BJf6OBP_SQaSBy7PJJUnBA-u5op_35sgACiKE"}
	//$webRequest.Content
	r.POST("/addProduct", service.AddProduct)
	//先登录获取令牌
	//$response = Invoke-WebRequest -Uri "http://localhost:8080/login" -Method POST -Body "useraccount=jiang&password=123"
	//$token = ($response.Content | ConvertFrom-Json).token
	//再添加商品
	//$headers = @{"Authorization"="Bearer $token"}
	//$body = @{userAccount="jiang"; productName="可乐"; productNumber=3; productPrice=2.50}
	//Invoke-RestMethod -Uri "http://localhost:8080/addProduct" -Method Post -Headers $headers -Body $body
	r.POST("/addBalance", service.AddBalance)
	//$headers = @{
	//   "Authorization" = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxhbiIsImV4cCI6MTcwNjExNTUxMiwiaXNzIjoi6JKL5Y2T54eDIn0.pFXOZ9HCX7zyodyWZIJL7I5l8vGRHCROPvcq8YebMjk"
	//}
	//
	//$body = @{
	//   "useraccount" = "lan"
	//   "money" = 100
	//} | ConvertTo-Json
	//
	//Invoke-WebRequest -Uri "http://localhost:8080/addBalance" -Method POST -Headers $headers -Body $body -ContentType "application/json"
	//
	r.POST("/makeCar", service.MakeCar)
	//$headers = @{
	//    "Authorization" = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxhbiIsImV4cCI6MTcwNjExNTUxMiwiaXNzIjoi6JKL5Y2T54eDIn0.pFXOZ9HCX7zyodyWZIJL7I5l8vGRHCROPvcq8YebMjk"
	//}
	//
	//Invoke-WebRequest -Uri "http://localhost:8080/makeCar" -Method POST -Headers $headers -ContentType "application/json"
	r.POST("/addCar", service.AddCar)
	//$headers = @{
	//    "Authorization" = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxhbiIsImV4cCI6MTcwNjExNTUxMiwiaXNzIjoi6JKL5Y2T54eDIn0.pFXOZ9HCX7zyodyWZIJL7I5l8vGRHCROPvcq8YebMjk"
	//}
	//
	//$body = @{
	//    "BuyAllProductID" = 1
	//    "BuyProductNumber" = 2
	//    "BuyUserAccount" = "lan"
	//} | ConvertTo-Json
	//
	//Invoke-WebRequest -Uri "http://localhost:8080/addCar" -Method POST -Headers $headers -Body $body -ContentType "application/json"
	r.POST("/lookCar", service.LookCar)
	//$headers = @{
	//   "Authorization" = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxhbiIsImV4cCI6MTcwNjExNTUxMiwiaXNzIjoi6JKL5Y2T54eDIn0.pFXOZ9HCX7zyodyWZIJL7I5l8vGRHCROPvcq8YebMjk"
	//   "Content-Type" = "application/json"
	//}
	//
	//$response = Invoke-WebRequest -Uri "http://localhost:8080/lookCar" -Method POST -Headers $headers
	//$carList = $response.Content | ConvertFrom-Json
	//$carList
	r.POST("/emptyCar", service.EmptyCar)
	//$headers = @{
	//    "Authorization" = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxhbiIsImV4cCI6MTcwNjExNTUxMiwiaXNzIjoi6JKL5Y2T54eDIn0.pFXOZ9HCX7zyodyWZIJL7I5l8vGRHCROPvcq8YebMjk"
	//    "Content-Type" = "application/json"
	//}
	//
	//$response = Invoke-WebRequest -Uri "http://localhost:8080/emptyCar" -Method POST -Headers $headers
	//$response.Content
	r.POST("/myProductList", service.MyProductList)
	//$headers = @{
	//    "Authorization" = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxhbiIsImV4cCI6MTcwNjEyMDA3NywiaXNzIjoi6JKL5Y2T54eDIn0.BVrYyjMRtjSG3A6Dolu9meQo0x1RqVT8iXCNitGDOKw"
	//}
	//
	//$body = @{
	//    "useraccount" = "lan"
	//}
	//
	//Invoke-WebRequest -Uri "http://localhost:8080/myProductList" -Method POST -Headers $headers -Body $body
	r.POST("/lookHostList", service.LookHostList)
	//$headers = @{
	//	   "Authorization" = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxhbiIsImV4cCI6MTcwNjExNTUxMiwiaXNzIjoi6JKL5Y2T54eDIn0.pFXOZ9HCX7zyodyWZIJL7I5l8vGRHCROPvcq8YebMjk"
	//	   "Content-Type" = "application/json"
	//	}
	//
	//	$response = Invoke-WebRequest -Uri "http://localhost:8080/lookHostList" -Method POST -Headers $headers
	//	$carList = $response.Content | ConvertFrom-Json
	//	$carList
	r.POST("/changeAccount", service.ChangeAccount)
	//$headers = @{
	//    "Authorization" = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxhbiIsImV4cCI6MTcwNjEyMDA3NywiaXNzIjoi6JKL5Y2T54eDIn0.BVrYyjMRtjSG3A6Dolu9meQo0x1RqVT8iXCNitGDOKw"
	//}
	//
	//$body = @{
	//    UserAccount = "lan"
	//    Changelocation = "Nickname"
	//    Changetext = "lanfanya"
	//}
	//
	//Invoke-RestMethod -Uri "http://localhost:8080/changeAccount" -Method Post -Headers $headers -Body ($body | ConvertTo-Json)
	err := r.Run(":" + port)
	if err != nil {
		return err
	}
	return err
}
