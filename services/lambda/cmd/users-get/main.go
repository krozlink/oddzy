package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type GetRequest struct {
	Name string `json:"name"`
}

type GetResponse struct {
	Message string `json:"message"`
}

func HandleRequest(ctx context.Context, request GetRequest) (GetResponse, error) {
	response := GetResponse{
		Message: fmt.Sprintf("Greetings %s!", request.Name),
	}

	return response, nil
}

func main() {
	lambda.Start(HandleRequest)
}
