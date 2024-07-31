package redisson

import (
	"github.com/moji-open-source/moji-chat-server/setup"
	"github.com/redis/go-redis/v9"
)

func NewClient(env *setup.Env) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
}

type Redisson struct {
	*redis.Client
}
