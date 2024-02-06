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

	r.POST("/login", service.Login)

	r.POST("/findProduct", service.FindProduct)

	r.Use(middleware.JWTAuthMiddleware()) //这里开始以后的调用令牌保护
	r.POST("/call", service.Call)
	// Invoke-WebRequest -Uri "http://localhost:8080/call" -Method POST -Headers @{"Authorization"="Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Ij8iLCJleHAiOjE3MDYwMDkwNzgsImlzcyI6IuiSi-WNk-eHgyJ9.GMi1FUgLHxNo8wqmgp6nfiKxhkQsOdErf-Pl9lEu4b0"}

	r.POST("/myself", service.Myself)

	r.POST("/refresh", service.Refresh)

	r.POST("/addProduct", service.AddProduct)

	r.POST("/addBalance", service.AddBalance)

	r.POST("/makeCar", service.MakeCar)

	r.POST("/addCar", service.AddCar)

	r.POST("/lookCar", service.LookCar)

	r.POST("/emptyCar", service.EmptyCar)

	r.POST("/myProductList", service.MyProductList)

	r.POST("/lookHostList", service.LookHostList)

	r.POST("/changeAccount", service.ChangeAccount)

	err := r.Run(":" + port)
	if err != nil {
		return err
	}
	return err
}
