package dao

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// 用map创建一个假数据库
var Database = map[string]string{
	"Eode":  "awzsex",
	"yuyue": "985632",
}

// 注册用户
func AddUser(username, password string) {
	Database[username] = password
}

// 检查用户名是否存在
func FindUser(username string) bool {
	if Database[username] == "" {
		return false
	}
	return true
}

// 查找输入用户名对应的密码
func FindPasswordFromUsername(username string) string {
	return Database[username]
}

// 响应反馈函数
func RespSuccess(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": message,
	})
}

func RespFail(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  500,
		"message": message,
	})
}
func Writedata(k map[string]string) {
	file, err := os.OpenFile("数据字典.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("打开文件出错:", err)
		return
	}
	defer file.Close()
	fmt.Println("成功打开文件:", file.Name())

	for key, value := range k {
		line := fmt.Sprintf("%s: %s\n", key, value)
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("写入文件出错:", err)
			return
		}
	}
	fmt.Println("数据已成功记录到文件中。")
}
