package main

import (
	"GoEode/lanshan/lanshan6/web/api"
)

func main() {

	api.InitRouter()
}

//如果想登录，可以在终端输入例如：
//Invoke-WebRequest -Uri "http://localhost:8088/login" -Method POST -Body @{username="Eode"; password="awzsex"}
//如果想注册，可以在终端输入例如：
//Invoke-WebRequest -Uri "http://localhost:8088/register" -Method POST -Body @{username="qian"; password="before"}
//如果想修改密码，可以在终端输入例如：
// Invoke-WebRequest -Uri "http://localhost:8088/changepassword" -Method POST -Body @{username="qian"; password="beforepeople"
