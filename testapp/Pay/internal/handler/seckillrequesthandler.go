package handler

import (
	"context"
	"dianshang/testapp/Pay/database"
	"dianshang/testapp/Pay/middleware"
	"fmt"
	"strconv"
	"time"

	"encoding/json"
	"errors"
	"net/http"

	"dianshang/testapp/Pay/internal/logic"
	"dianshang/testapp/Pay/internal/svc"
	"dianshang/testapp/Pay/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type SeckillRequestMessage struct {
	ProductID   int    `json:"productId"`
	BuyQuantity int    `json:"buyQuantity"`
	UserName    string `json:"userName"`
}

// SeckillRequestHandler 申请秒杀请求
func SeckillRequestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SeckillRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 从哈希中获取商品信息
		result, err := database.ShopRedis1DB.HGetAll(context.Background(), strconv.Itoa(req.ProductId)).Result()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 获取商品的库存数量
		stock, err := strconv.Atoi(result["stock"])
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 检查库存是否足够
		if stock < req.BuyQuantity {
			httpx.ErrorCtx(r.Context(), w, errors.New("库存不足"))
			return
		}

		// 锁的键
		lockKey := fmt.Sprintf("product_%d", req.ProductId)

		// 获取锁
		lockValue, err := database.RedisLock(database.ShopRedis1DB, lockKey, 5*time.Second)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 减少库存原子操作
		err = database.ShopRedis1DB.HIncrBy(context.Background(), strconv.Itoa(req.ProductId), "stock", int64(-req.BuyQuantity)).Err()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 释放锁
		err = database.RedisUnlock(database.ShopRedis1DB, lockKey, lockValue)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 从上下文中获取 JWT 令牌
		username, ok := r.Context().Value(middleware.ContextKeyUsername).(string)
		if !ok {
			httpx.ErrorCtx(r.Context(), w, errors.New("username is missing from context"))
			return
		}

		// 创建一个 CreateOrderMessage 结构体实例
		msg := SeckillRequestMessage{
			ProductID:   req.ProductId,
			BuyQuantity: req.BuyQuantity,
			UserName:    username,
		}

		// 将 msg 转换为 JSON
		reqData, err := json.Marshal(msg)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 发送消息到 Kafka
		err = svcCtx.KafkaClient.SendMessage("SeckillRequest", reqData)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 将 username 放入上下文中
		ctx := context.WithValue(r.Context(), "UserName", username)

		l := logic.NewSeckillRequestLogic(ctx, svcCtx)
		resp, err := l.SeckillRequest(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
