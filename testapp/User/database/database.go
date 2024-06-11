package database

import "github.com/go-redis/redis/v8"

var UserRedis2DB = redis.NewClient(&redis.Options{
	Addr:     "localhost:26379",
	Password: "123awzsex",
	DB:       0, // use default DB
})
