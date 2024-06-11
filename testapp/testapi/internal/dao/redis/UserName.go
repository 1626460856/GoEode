package redis

import (
	"context"
	"dianshang/testapp/testapi/global"
	"errors"
	"fmt"
)

func CreateUserNameSet() {
	ctx := context.Background()

	// 使用 SADD 命令添加元素到 "UserName" 集合
	err := global.UserRedis1DB.SAdd(ctx, "UserName", "test").Err()
	if err != nil {
		global.Logger.Error("failed to add to set: %v" + err.Error())
		fmt.Errorf("failed to add to set: %v", err)
		return
	}
	fmt.Println("成功把test用户名添加到UserNameSet")
	return
}
func AddToSet(setName string, value string) error {
	ctx := context.Background()
	// 使用 SAdd 命令添加元素到集合，它会返回添加的元素数量
	added, err := global.UserRedis1DB.SAdd(ctx, setName, value).Result()
	if err != nil {
		return err
	}

	// 如果添加的元素数量为0，说明元素已经存在于集合中
	if added == 0 {
		global.Logger.Error("元素" + value + "在集合" + setName + "中已经存在")
		return errors.New("元素在集合中已经存在")
	}

	return nil
}
