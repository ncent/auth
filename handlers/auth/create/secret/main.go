package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	createSecretHelper "gitlab.com/ncent/auth/services/auth/create/secret"
)

type response struct {
	Secret string `json:"secret"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	resp := events.APIGatewayProxyResponse{
		Headers: make(map[string]string),
	}
	resp.Headers["Access-Control-Allow-Origin"] = "*"

	// Get public key parameter
	publicKey, ok := request.QueryStringParameters["publicKey"]

	// Return Bad Request if Public Key is not passed
	if !ok || len(publicKey) == 0 {
		resp.StatusCode = http.StatusBadRequest
		return resp, nil
	}

	// Get Secret
	secret, err := createSecretHelper.GetSecret(publicKey)
	if err != nil {
		log.Printf("Failed to retrive secret %+v", err)
		return resp, err
	}

	// Create response with secret
	response, err := json.Marshal(response{Secret: *secret})
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
