package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// GetRedisValue 获取键值对
// ctx: context.Context 类型，提供请求的上下文信息，通常用于控制超时或取消信号
// key: string 类型，表示要从 Redis 中获取的键名。
// 这个函数从 Redis 中获取指定键的值。如果获取过程中发生错误，则返回空字符串和错误信息；否则，返回获取到的值和 nil（表示没有错误）。
func GetRedisValue(ctx context.Context, key string) (string, error) {
	GetKey := Rdb.Get(ctx, key)
	if GetKey.Err() != nil {
		return "", GetKey.Err()
	}
	return GetKey.Val(), nil
}

// SetRedisValue 存储键值对
// ctx: context.Context 类型，提供请求的上下文信息。
// key: string 类型，表示要设置的键名。
// value: string 类型，表示要设置的值。
// expiration: time.Duration 类型，表示键值对的过期时间。
func SetRedisValue(ctx context.Context, key string, value string, expiration time.Duration) error {
	SetKV := Rdb.Set(ctx, key, value, expiration)
	return SetKV.Err()
}

// RedisSet  集合成员结构体 集合名称为key 其中包含多个member

type RedisSet struct {
	key     string
	member  string
	Conn    *redis.Client
	Context context.Context
}

func NewRedisSet(context context.Context, key string, member string, Conn *redis.Client) *RedisSet {
	return &RedisSet{
		key:     key,
		member:  member,
		Conn:    Conn,
		Context: context,
	}
}

// NewSet 往指定名称的集合中添加一个空元素，以此来创建集合
func NewSet(key string) {
	rs := NewRedisSet(context.Background(), key, "", Rdb)
	_, err := rs.Conn.SAdd(rs.Context, rs.key, rs.member).Result()
	if err != nil {
		fmt.Println(err)
	}
	if err == nil {
		fmt.Println(key + "集合创建成功")
	}
}

// AddSetMember 往指定集合中添加一个成员
func AddSetMember(key string, member string) {
	rs := NewRedisSet(context.Background(), key, member, Rdb)
	_, err := rs.Conn.SAdd(rs.Context, rs.key, rs.member).Result()
	if err != nil {
		fmt.Println(err)
	}
	if err == nil {
		fmt.Println("成功往" + key + "集合添加成员：" + member)
	}
}
