package handler

import (
	"context"
	"dianshang/testapp/User/database"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"dianshang/testapp/User/internal/logic"
	"dianshang/testapp/User/internal/svc"
	"dianshang/testapp/User/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddUserToSet(UserRedis2DB *redis.Client, setName string, value string) error {
	ctx := context.Background()
	// 使用 SAdd 命令添加元素到集合，它会返回添加的元素数量
	fmt.Println(UserRedis2DB)
	added, err := UserRedis2DB.SAdd(ctx, setName, value).Result()
	if err != nil {
		return err
	}

	// 如果添加的元素数量为0，说明元素已经存在于集合中
	if added == 0 {
		logx.Errorf("用户 %s 在用户集 %s 中已经存在", value, setName)
		return errors.New("用户" + value + "在用户集合" + setName + "中已经存在")
	}

	return nil
}
func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		// 检查用户名是否已存在
		err := AddUserToSet(database.UserRedis2DB, "UserName", req.UserName)
		fmt.Println("AddToSet err:", err)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		// 将请求数据转换为 JSON
		reqData, err := json.Marshal(req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 发送消息到 Kafka
		err = svcCtx.KafkaClient.SendMessage("RegisterReq", reqData)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 调用业务逻辑处理
		//创建一个新的 RegisterLogic 实例 l，并调用其 Register 方法来处理业务逻辑
		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
