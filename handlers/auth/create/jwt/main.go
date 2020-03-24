package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	createJWTHelper "gitlab.com/ncent/auth/services/auth/create/jwt"
)

type response struct {
	JWT string `json:"JWT"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	resp := events.APIGatewayProxyResponse{
		Headers: make(map[string]string),
	}
	resp.Headers["Access-Control-Allow-Origin"] = "*"

	// Get email and public key parameters
	userEmail, okEmail := request.QueryStringParameters["email"]
	publicKey, ok := request.QueryStringParameters["publicKey"]
	signature, okSign := request.QueryStringParameters["signature"]

	// Return Bad Request if Public Key is not passed
	if !okSign || !ok || !okEmail || len(publicKey) == 0 || len(signature) == 0 || len(userEmail) == 0 {
		resp.StatusCode = http.StatusBadRequest
		return resp, nil
	}

	// Get Secret
	jwt, err := createJWTHelper.GetJWT(userEmail, publicKey, signature)
	if err != nil {
		log.Printf("Failed to retrive JWT %+v", err)
		return resp, err
	}

	// Create response with secret
	response, err := json.Marshal(response{JWT: *jwt})
	if err != nil {
		return resp, err
	}

	resp.StatusCode = http.StatusOK
	resp.Body = string(response)

	return resp, nil
}

func main() {
	lambda.Start(handler)
}
