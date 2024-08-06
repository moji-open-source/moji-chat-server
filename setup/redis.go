package setup

import "github.com/redis/go-redis/v9"

func NewRedlisClient(env *Env) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:       env.RedisAddr,
		Password:   env.RedisPassword, // 没有密码，默认值
		DB:         env.RedisDatabase, // 默认DB 0
		Username:   env.RedisUsername,
		PoolSize:   env.RedisPollSize,
		ClientName: env.RedisClientName,
	})
}
