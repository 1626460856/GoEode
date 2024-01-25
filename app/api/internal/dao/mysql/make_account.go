package mysql

import (
	"database/sql"
	"dianshang/app/api/global"
	"fmt"
	"go.uber.org/zap"
)

func MakeAccountList(db *sql.DB) {
	query := `
        CREATE TABLE IF NOT EXISTS AccountList (
            ID INT AUTO_INCREMENT PRIMARY KEY,
            UserAccount VARCHAR(50) NOT NULL,
            Password VARCHAR(50) NOT NULL,
            Nickname VARCHAR(100), 
            Identity VARCHAR(20) NOT NULL, 
            Balance DOUBLE NOT NULL DEFAULT 0.0          
        );
    `

	_, err := db.Exec(query)
	if err != nil {
		global.Logger.Fatal("创建表失败 ," + err.Error())
		return
	}
	global.Logger.Info("已成功创建 AccountList 表")
	return
}

func MakeAllProductList(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS AllProductList (
			AllProductID INT AUTO_INCREMENT PRIMARY KEY,
			ProductListName VARCHAR(50) NOT NULL,
			ProductID INT,
			ProductName VARCHAR(100) NOT NULL,
			ProductNumber INT NOT NULL,
			ProductPrice DOUBLE NOT NULL
		);
	`

	_, err := db.Exec(query)
	if err != nil {
		global.Logger.Fatal("创建表失败," + err.Error())
		return
	}
	global.Logger.Info("已成功创建 AllProductList 表")
}
func AddAccount(db *sql.DB, useraccount, password, identity, nickname string) error {
	// 检查用户名是否已存在
	var count int
	query := "SELECT COUNT(*) FROM AccountList WHERE UserAccount = ?"
	err := db.QueryRow(query, useraccount).Scan(&count)
	if err != nil {
		global.Logger.Error("输入读取错误", zap.Error(err))
		return fmt.Errorf("输入读取错误", err.Error())
	}
	if count > 0 {
		global.Logger.Error("该用户名已被占用")
		return fmt.Errorf("该用户名已被占用", err.Error())
	}

	// 执行插入操作
	insertQuery := "INSERT INTO AccountList (UserAccount, Password, Identity, Nickname) VALUES (?, ?, ?, ?)"
	_, err = db.Exec(insertQuery, useraccount, password, identity, nickname)
	if err != nil {
		global.Logger.Error("无法插入用户数据", zap.Error(err))
		return fmt.Errorf("无法插入用户数据", err.Error())
	}
	// 创建个人商品库
	err = MakeProductList(db, useraccount)
	if err != nil {
		global.Logger.Error("无法在创建新账户的时候创建个人商品库", zap.Error(err))
		return fmt.Errorf("无法在创建新账户的时候创建个人商品库", err.Error())
	}
	global.Logger.Info("成功在注册新账户的时候创建个人商品库")
	return nil
}
func MakeProductList(db *sql.DB, useraccount string) error {
	query := `
		CREATE TABLE IF NOT EXISTS ` + useraccount + ` (
		    ProductID INT AUTO_INCREMENT PRIMARY KEY,
		    ProductName VARCHAR(100) NOT NULL,
		    ProductNumber INT NOT NULL,
		    ProductPrice DOUBLE NOT NULL
		);
	`
	_, err := db.Exec(query)
	if err != nil {
		global.Logger.Fatal("create table failed," + err.Error())
		return fmt.Errorf("无法创建个人商品库")
	}
	global.Logger.Info(useraccount + " 商品库创建成功")
	return nil
}

// DropList 删库跑路函数
func DropList(db *sql.DB, ListName string) {
	query := "DROP TABLE IF EXISTS " + ListName
	_, err := db.Exec(query)
	if err != nil {
		global.Logger.Fatal("drop table failed: " + err.Error())
	}
	global.Logger.Info(ListName + " 表格删除成功")
	//global.Logger.Info("成功删库跑路")
}
