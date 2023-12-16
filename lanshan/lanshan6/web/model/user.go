package model

import "github.com/dgrijalva/jwt-go"

// 用户的结构体为用户名和密码
type User struct {
	Username string
	Password string
}
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
