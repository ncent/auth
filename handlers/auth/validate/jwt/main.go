package main

import (
	"errors"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	jwtService "github.com/dgrijalva/jwt-go"
	jwtValidateService "gitlab.com/ncent/auth/services/auth/validate/jwt"
	policyService "gitlab.com/ncent/auth/services/aws/policy"
)

func handler(request events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	authorizationToken := request.AuthorizationToken
	tokenSlice := strings.Split(authorizationToken, " ")
	var bearerToken string
	if len(tokenSlice) > 1 {
		bearerToken = tokenSlice[len(tokenSlice)-1]
	}

	token, err := jwtValidateService.GetJWTToken(bearerToken)
	if (token != nil && !token.Valid) || err != nil {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	claims, ok := token.Claims.(jwtService.MapClaims)
	if !ok {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	response := policyService.GeneratePolicy(claims["publicKey"].(string), "Allow", request.MethodArn)
	return response, nil
}

func main() {
	lambda.Start(handler)
}
