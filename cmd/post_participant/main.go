package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"hello-world/internal/infrastructure/config"
	"hello-world/internal/infrastructure/go/util/err"
	"hello-world/internal/interfaces/lambda/participant"
)

func main() {
	cfg := err.PanicIfError1(config.NewConfig())
	handler := participant.NewLambdaHandler(cfg)

	lambda.Start(handler.HandleRequest)
}
