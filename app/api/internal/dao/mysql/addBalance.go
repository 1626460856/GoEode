package mysql

import (
	"database/sql"
	"dianshang/app/api/global"
)

func AddBalance(db *sql.DB, UserAccount string, money float64) error {
	_, err := db.Exec("UPDATE AccountList SET Balance = Balance + ? WHERE UserAccount = ?", money, UserAccount)
	if err != nil {
		global.Logger.Fatal("更新账户余额失败," + err.Error())
		return err
	}
	return nil
}
