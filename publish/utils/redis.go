package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()
var RdbVideoId *redis.Client

// InitRedis 初始化Redis连接。
func InitRedis() {

	RdbVideoId = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       1, // 视频信息存入 DB1.
	})
}
