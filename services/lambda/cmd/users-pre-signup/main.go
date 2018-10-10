package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
}

type Response struct {
	AutoConfirmUser string `json:"autoConfirmUser"`
	AutoVerifyPhone string `json:"autoVerifyUser"`
	AutoVerifyEmail string `json:"autoVerifyEmail"`
}

func HandleRequest(ctx context.Context, request Request) (Response, error) {
	response := Response{
		AutoConfirmUser: "true",
		AutoVerifyEmail: "false",
		AutoVerifyPhone: "false",
	}

	return response, nil
}

func main() {
	lambda.Start(HandleRequest)
}
