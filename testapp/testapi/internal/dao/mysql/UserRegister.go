package mysql

import (
	"dianshang/testapp/testapi/global"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Id           int
	UserName     string
	PassWord     string
	UserNick     string
	UserIdentity string
	balance      float64 = 0
)

func CreateRegisterUsersTable() {
	// SQL语句
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS RegisterUsers (
		id INT AUTO_INCREMENT,
		username VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		usernick VARCHAR(255) NOT NULL,
		userIdentity VARCHAR(255) NOT NULL,
		balance FLOAT DEFAULT 0,
		PRIMARY KEY (id)
	);`

	// 执行SQL语句
	_, err := global.UserMysqlDB.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	fmt.Println("Table created successfully")
}
