package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()
var RdbUserVideo *redis.Client
var RdbVideoId *redis.Client
var RdbUserId *redis.Client

func InitRedis() {
	RdbUserVideo = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       2, // 是否点赞信息存入 DB2.
	})
	RdbVideoId = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       1, // 视频信息存入 DB1.
	})
	RdbUserId = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0, // 用户信息存入 DB0.
	})
}
