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
		"quizId":"fcf7890f-9c72-46d3-931e-34494307be37",
		"requiredQuestionsAnswered": ["f5b70d7d-3461-4cf7-978d-2b0caf77db1e"]
	},
   "type": "%s"
}
`, data.StartQuizCommandType)

var eventBody2 = fmt.Sprintf(`
{
   "createdAt":"2023-11-17T04:55:24.059Z",
   "data": {
		"quizId":"fcf7890f-9c72-46d3-931e-34494307be37",
		"questionId":"f5b70d7d-3461-4cf7-978d-2b0caf77db1e",
		"answerId": "1eb5a02f-562a-409d-8822-a240e0886485"
	},
   "type": "%s"
}
`, data.SelectAnswerCommandType)

var eventBody3 = fmt.Sprintf(`
{
   "createdAt":"2023-11-17T04:55:24.059Z",
   "data": {
		"quizId":"fcf7890f-9c72-46d3-931e-34494307be37"
	},
   "type": "%s"
}
`, data.FinishQuizCommandType)

func TestPutParticipantLambda_Returns200(t *testing.T) {

	environmentCreator := local.NewEnvironmentCreator(config.Test)

	requestBodys := []string{
		eventBody,
		eventBody2,
		eventBody3,
	}

	handleRequestFn := participant.NewLambdaHandler(environmentCreator.Cfg).HandleRequest
	for _, requestBody := range requestBodys {
		handlerResponse := environmentCreator.ExecuteLambdaHandlerWithPostBody(
			handleRequestFn,
			requestBody,
		)
		if handlerResponse.StatusCode != 200 {
			t.Fatalf("lambda did not succeed, status code: %v, %v", handlerResponse.StatusCode, handlerResponse.Body)
		}
	}

}

func TestPutParticipantLambda_InvalidQuizCompletion_Returns500(t *testing.T) {

	environmentCreator := local.NewEnvironmentCreator(config.Test)

	handleRequestFn := participant.NewLambdaHandler(environmentCreator.Cfg).HandleRequest

	environmentCreator.ExecuteLambdaHandlerWithPostBody(
		handleRequestFn,
		eventBody,
	)

	handlerResponse3 := environmentCreator.ExecuteLambdaHandlerWithPostBody(
		handleRequestFn,
		eventBody3,
	)
	if handlerResponse3.StatusCode != 500 {
		t.Fatalf("lambda return code is not 500 although invalid quiz completion case: %v, %v", handlerResponse3.StatusCode, handlerResponse3.Body)
	}

}
