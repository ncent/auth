package jwt

import (
	"errors"
	"fmt"
	"os"

	jwtService "github.com/dgrijalva/jwt-go"
)

func GetJWTToken(jwtString string) (*jwtService.Token, error) {

	var hmacSecret, ok = os.LookupEnv("JWT_SECRET")
	if !ok {
		fmt.Println("JWT_SECRET is required")
		return nil, errors.New("JWT_SECRET is required")
	}

	token, err := jwtService.Parse(jwtString, func(token *jwtService.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtService.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(hmacSecret), nil
	})
	if err != nil {
		fmt.Println("Could not parse JWT token")
		return nil, err
	}

	return token, nil
}
