package main_test

import (
	"fmt"
	"learn-to-code/internal/domain/command/data"
	"learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/local"
	"learn-to-code/internal/interfaces/lambda/participant"
	"testing"
)

var eventBody = fmt.Sprintf(`
{
   "createdAt":"2023-11-17T04:55:24.059Z",
   "data": {
		"quizId":"fcf7890f-9c72-46d3-931e-34494307be37"
	},
   "type": "%s"
}
`, data.StartQuizCommandType)

func TestPutParticipantLambda_Returns200(t *testing.T) {

	environmentCreator := local.NewEnvironmentCreator(config.Test)
	handlerResponse := environmentCreator.ExecuteLambdaHandlerWithPostBody(
		participant.NewLambdaHandler(environmentCreator.Cfg).HandleRequest,
		eventBody,
	)

	if handlerResponse.StatusCode != 200 {
		t.Fatalf("lambda did not succeed, status code: %v, %v", handlerResponse.StatusCode, handlerResponse.Body)
	}
}
