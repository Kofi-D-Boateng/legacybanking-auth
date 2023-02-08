package utils

import (
	"context"

	redis "github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func ConnectRedis(addr string, password string, database int) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       database,
	})
	err := RedisClient.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
}
