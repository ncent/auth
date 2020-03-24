package jwt

import (
	"errors"
	"fmt"
	"log"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	crypto "gitlab.com/ncent/auth/services/crypto"
	RedisService "gitlab.com/ncent/auth/services/redis"
)

func GetJWT(userEmail string, publicKey string, signature string) (*string, error) {

	if len(userEmail) == 0 {
		return nil, errors.New("Empty User Email")
	}

	if len(publicKey) == 0 {
		return nil, errors.New("Empty Public Key")
	}

	if len(signature) == 0 {
		return nil, errors.New("Empty Signature")
	}

	secret, err := retrieveSecret(publicKey)
	if err != nil {
		return nil, err
	}

	valid, err := crypto.IsValidSignature(publicKey, signature, *secret)
	if err != nil {
		log.Printf("Failed to validate signature %+v", err)
		return nil, err
	}

	if !valid {
		return nil, errors.New("Invalid Signature")
	}

	tokenString, err := CreateSignedToken(userEmail, publicKey)
	if err != nil {
		return nil, err
	}

	return tokenString, nil
}

func retrieveSecret(publicKey string) (*string, error) {
	redisService := RedisService.New()

	secret, err := redisService.Client.Get(publicKey).Result()
	if err != nil && err != redis.Nil {
		log.Printf("Failed to get key %+v", err)
		return nil, err
	}
	if secret == "" {
		return nil, errors.New("Secret not found")
	}

	return &secret, nil
}

func CreateSignedToken(userEmail string, publicKey string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"publicKey": publicKey,
		"email":     userEmail,
	})

	var hmacSecret, ok = os.LookupEnv("JWT_SECRET")
	if !ok {
		fmt.Println("JWT_SECRET is required")
		return nil, errors.New("JWT_SECRET is required")
	}

	tokenString, err := token.SignedString([]byte(hmacSecret))
	if err != nil {
		log.Printf("Failed to get token string %+v", err)
		return nil, err
	}

	return &tokenString, nil
}
