package main

import (
	"learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/go/util/err"
	"learn-to-code/internal/infrastructure/service"
	"learn-to-code/internal/interfaces/lambda/participant/quiz"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	cfg := err.PanicIfError1(config.NewConfig())

	getParticipantQuizOverviewHandler := quiz.NewGetParticipantQuizAttemptDetailHandler(cfg, service.RegistryOverride{})

	lambda.Start(getParticipantQuizOverviewHandler.HandleRequest)
}
