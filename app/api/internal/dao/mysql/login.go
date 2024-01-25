package mysql

import (
	"database/sql"
	"dianshang/app/api/global"
	"dianshang/app/api/internal/middleware"
	"dianshang/app/api/internal/model"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
	"time"
)

func FindUser(db *sql.DB, UserAccount string, ListName string) {
	query := "SELECT Id, UserAccount, Password, Nickname, Identity, Balance FROM " + ListName + " WHERE UserAccount = ?"
	row := db.QueryRow(query, UserAccount)
	var (
		id          int
		userAccount string
		password    string
		nickname    sql.NullString
		identity    string
		balance     float64
	)
	err := row.Scan(&id, &userAccount, &password, &nickname, &identity, &balance)
	if err != nil {
		global.Logger.Error("Failed to find user", zap.Error(err))
		return
	}
	fmt.Printf("User found:\nID: %v\nUserAccount: %s\nPassword: %s\nNickname: %s\nIdentity: %s\nBalance: %f\n",
		id, userAccount, password, nickname.String, identity, balance)
}
func CoutFindUser(db *sql.DB, UserAccount string, ListName string) (model.Account, error) {
	query := "SELECT Id, UserAccount, Password, Nickname, Identity, Balance FROM " + ListName + " WHERE UserAccount = ?"
	row := db.QueryRow(query, UserAccount)
	var account model.Account
	err := row.Scan(&account.ID, &account.UserAccount, &account.Password, &account.Nickname, &account.Identity, &account.Balance)
	if err != nil {
		global.Logger.Error("无法扫描行: " + err.Error())
		return account, errors.New("无法扫描行")
	}

	return account, nil
}

func FindUserAccount(db *sql.DB, UserAccount string) (string, error) {
	query := "SELECT Password FROM AccountList WHERE UserAccount = ?"
	var password string
	err := db.QueryRow(query, UserAccount).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("用户名不存在")
		}
		return "", err
	}
	return password, nil
}
func Login(db *sql.DB, UserAccount string, Password string) (string, error) {
	// 在数据库中验证用户名和密码
	password, err := FindUserAccount(db, UserAccount)
	if err != nil {
		global.Logger.Error("登录失败，该用户不存在", zap.Error(err))
		return "登录失败，该用户不存在", err
	}
	if password == Password {

		// 登录成功，创建JWT令牌
		claims := model.MyClaims{
			Username: UserAccount,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
				Issuer:    "蒋卓燃",
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(middleware.Secret)
		if err != nil {
			return "", err
		}
		return tokenString, nil

	}
	return "", errors.New("密码错误")
}
