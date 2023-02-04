package cache

import (
	"os"

	"github.com/go-redis/redis"
)

func GetRedisConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URI"),
		Password: "",
		DB:       0,
	})
}


