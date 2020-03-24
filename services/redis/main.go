package redis

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

type RedisService struct {
	Client *redis.Client
}

func New() RedisService {
	var REDIS_ENDPOINT, ok = os.LookupEnv("REDIS_ENDPOINT")
	if !ok {
		fmt.Println("REDIS_ENDPOINT is required")
	}
	client := redis.NewClient(&redis.Options{
		Addr: REDIS_ENDPOINT,
	})
	return RedisService{client}
}
