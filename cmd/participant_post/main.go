package main

import (
	"learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/go/util/err"
	"learn-to-code/internal/infrastructure/service"
	"learn-to-code/internal/interfaces/lambda/participant"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	cfg := err.PanicIfError1(config.NewConfig())

	handler := participant.NewPostParticipantCommandHandler(cfg, service.RegistryOverride{})

	lambda.Start(handler.HandleRequest)
}
