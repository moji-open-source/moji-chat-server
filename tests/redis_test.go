package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type RedisKey = string

const (
	Authentication RedisKey = "Authentication:mojichat:login"
)

func TestRedis(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

	ctx := context.Background()
	err := rdb.Set(ctx, Authentication+":1", uuid.New().String(), -1).Err()
	if err != nil {
		fmt.Println("err = ", err)
	}
}
