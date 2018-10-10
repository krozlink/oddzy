package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.CognitoEventUserPoolsPreSignup) (events.CognitoEventUserPoolsPreSignup, error) {
	event.Response.AutoConfirmUser = true
	event.Response.AutoVerifyEmail = false
	event.Response.AutoVerifyPhone = false

	return event, nil
}

func main() {
	lambda.Start(handler)
}
