package setup

import "github.com/redis/go-redis/v9"

func NewRedlisClient(env *Env) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
}
