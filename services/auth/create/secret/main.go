package secret

import (
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/teris-io/shortid"
	RedisService "gitlab.com/ncent/auth/services/redis"
)

const ONE_MINUTE_EXPIRATION_TIME = time.Duration(1) * time.Minute

func GetSecret(publicKey string) (*string, error) {

	if len(publicKey) == 0 {
		return nil, errors.New("Invalid Public Key")
	}
	redisService := RedisService.New()

	secret, err := redisService.Client.Get(publicKey).Result()
	if err != nil && err != redis.Nil {
		log.Printf("Failed to get key %+v", err)
		return nil, err
	}
	if secret == "" {
		secret = shortid.MustGenerate()
		redisService.Client.Set(publicKey, secret, ONE_MINUTE_EXPIRATION_TIME)
	}

	return &secret, nil
}
