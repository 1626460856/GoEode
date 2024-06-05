package model

import "github.com/dgrijalva/jwt-go"

type Account struct {
	UserAccount string  //用户账号名称
	Password    string  //用户密码
	ID          int     //从1开始计，第一个用户ID是1，自增键
	Nickname    string  //用户昵称
	Identity    string  //boss 或者 customer
	Balance     float64 //账户余额
}

type ProductList struct { //表格名称与UserAccount保持一致
	ProductID     int     //每个用户的商品表从1开始自增键
	ProductName   string  //商品名称
	ProductNumber int     //商品数量
	ProductPrice  float64 //商品价格
}
type AllProductList struct {
	AllProductID    int    //从1开始自增键
	ProductListName string //商品库名称
	ProductID       int
	ProductName     string  //商品名称
	ProductNumber   int     //商品数量
	ProductPrice    float64 //商品价格

}
type Car struct {
	Id              int    //自增键？
	AllProductID    int    //在所有商品表中的目录id
	ProductListName string //商品库名称
	ProductID       int
	ProductName     string  //商品名称
	ProductNumber   int     //商品数量
	ProductPrice    float64 //商品价格
}
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type HostList struct {
	ProductName string
	SoldNumber  int
	Price       float64
}
