package redis

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	redisAddr := os.Getenv("REDIS_URL")
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	return client
}
