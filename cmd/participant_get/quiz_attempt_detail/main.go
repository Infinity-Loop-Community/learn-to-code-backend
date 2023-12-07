package main

import (
	"learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/go/util/err"
	"learn-to-code/internal/interfaces/lambda/participant/quiz"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	cfg := err.PanicIfError1(config.NewConfig())

	getParticipantQuizOverviewHandler := quiz.NewGetParticipantQuizAttemptDetailHandler(cfg)

	lambda.Start(getParticipantQuizOverviewHandler.HandleRequest)
}
