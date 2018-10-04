package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type UserRequest struct {
	Name string `json:"name"`
}

type UserResponse struct {
	Message string `json:"message"`
}

func HandleRequest(ctx context.Context, request UserRequest) (UserResponse, error) {
	response := UserResponse{
		Message: fmt.Sprintf("Hello %s!", request.Name),
	}

	return response, nil
}

func main() {
	lambda.Start(HandleRequest)
}
