package utils

import redis "github.com/go-redis/redis/v8"

var RedisClient *redis.Client

func ConnectRedis(addr string, password string, database int) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       database,
	})
}
