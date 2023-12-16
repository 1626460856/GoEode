package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func findUser(db *sql.DB, user string) {
	query := fmt.Sprintf("SELECT schema_name FROM INFORMATION_SCHEMA.SCHEMATA WHERE schema_name LIKE '%%%s%%'", user)
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	var dbname string
	found := false
	for rows.Next() {
		err := rows.Scan(&dbname)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("查询到了匹配用户：", dbname)
		found = true
	}
	if !found {
		fmt.Println("未查询到任何用户")
	}
}

func combineStrings(A string, B string) string {
	return A + B
}

func addUser(db *sql.DB, newUsername string, newQQ string) {
	newUser := combineStrings(newUsername, newQQ)
	rows, err := db.Query("SHOW DATABASES")
	if err != nil {
		panic(err.Error())
	}
	var dbname string
	found := false
	for rows.Next() {
		err := rows.Scan(&dbname)
		if err != nil {
			panic(err.Error())
		}
		if newUser == dbname {
			found = true
			fmt.Println("该用户名已被占用：", dbname)
		}
	}
	if !found {
		_, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + newUser + " DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci")
		if err != nil {
			panic(err.Error())
		}

		// 使用新创建的数据库
		_, err = db.Exec("USE " + newUser)
		if err != nil {
			panic(err.Error())
		}

		// 创建表
		_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (ID INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255), wenzi VARCHAR(255))")
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("添加了账号：", newUser)
	}
}

func addFriend(db *sql.DB, findname string, username string) {
	findUser(db, findname)
	fmt.Println("是否继续添加好友？如果是，请输入上述查询结果完整名称：")
	var k string
	_, err := fmt.Scanf("%s", &k)
	if err != nil {
		panic(err)
	}
	if k == " " {
		return
	}
	if k != " " {
		err := write(db, k, username)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("成功添加好友：", k)
		return
	}
}

func write(db *sql.DB, name string, username string) error {
	insertStmt := fmt.Sprintf("INSERT INTO %s (name, wenzi) VALUES (?, ?)", username)
	_, err := db.Exec(insertStmt, name, "")
	if err != nil {
		return err
	}
	return nil
}

func showAllUsers(db *sql.DB) {
	rows, err := db.Query("SHOW DATABASES")
	if err != nil {
		panic(err.Error())
	}

	var dbname string
	for rows.Next() {
		err := rows.Scan(&dbname)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(dbname)
	}

	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}
}

func showUserData(db *sql.DB, username string) {
	// 切换到指定的数据库
	_, err := db.Exec("USE " + username)
	if err != nil {
		panic(err.Error())
	}

	// 查询用户数据
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}

	var (
		ID    int
		name  string
		wenzi string
	)

	for rows.Next() {
		err := rows.Scan(&ID, &name, &wenzi)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("ID:", ID, "Name:", name, "Wenzi:", wenzi)
	}

	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/mysql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MySQL database!")

	// 在此处添加您的数据库操作代码j
	addUser(db, "蒋卓燃", "1626460856")
	showAllUsers(db)
	showUserData(db, "蒋卓燃1626460856")
	addUser(db, "甲", "11111")
	showAllUsers(db)
	addFriend(db, "甲", "蒋卓燃1626460856")
	showUserData(db, "蒋卓燃1626460856")
}
