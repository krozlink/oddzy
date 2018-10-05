package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type AddRequest struct {
	Name string `json:"name"`
}

type AddResponse struct {
	Message string `json:"message"`
}

func HandleRequest(ctx context.Context, request AddRequest) (AddResponse, error) {
	response := AddResponse{
		Message: fmt.Sprintf("Greetings %s!", request.Name),
	}

	return response, nil
}

func main() {
	lambda.Start(HandleRequest)
}
