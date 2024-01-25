package mysql

import (
	"database/sql"
	"dianshang/app/api/global"
	"errors"
)

func ChangeAccount(db *sql.DB, UserAccount string, Changelocation string, Changetext string) error {
	var query string
	switch Changelocation {
	case "Password":
		query = "UPDATE AccountList SET Password=? WHERE UserAccount=?"
	case "Identity":
		if Changetext != "boss" && Changetext != "customer" {
			global.Logger.Error("无效的身份修改")
			return errors.New("无效的身份修改")
		}
		query = "UPDATE AccountList SET Identity=? WHERE UserAccount=?"
	case "Nickname":
		query = "UPDATE AccountList SET Nickname=? WHERE UserAccount=?"
	default:
		global.Logger.Error("修改位置传入错误")
		return errors.New("修改位置传入错误")
	}

	_, err := db.Exec(query, Changetext, UserAccount)
	if err != nil {
		global.Logger.Error("执行更新操作时出错: " + err.Error())
		return err
	}

	global.Logger.Info("成功将" + UserAccount + "的" + Changelocation + "修改为：" + Changetext)
	return nil
}
