package utils

import redis "github.com/go-redis/redis/v8"


var redisClient *redis.Client

func ConnectClient(addr string, password string, database int){
	redisClient = redis.NewClient(&redis.Options{
		Addr: addr,
		Password: password,
		DB: database,
	})
}

