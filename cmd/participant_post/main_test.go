package main_test

import (
	"fmt"
	"learn-to-code/internal/domain/command"
	"learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/local"
	"learn-to-code/internal/interfaces/lambda/participant"
	"testing"
)

var startQuizCommand = fmt.Sprintf(`
{
   "createdAt":"2023-11-17T04:55:24.059Z",
   "data": {
		"quizId":"fcf7890f-9c72-46d3-931e-34494307be37",
		"requiredQuestionsAnswered": ["f5b70d7d-3461-4cf7-978d-2b0caf77db1e"]
	},
   "type": "%s"
}
`, command.StartQuizCommandType)

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
`, command.SelectAnswerCommandType)

var finishQuizComand = fmt.Sprintf(`
{
   "createdAt":"2023-11-17T04:55:24.059Z",
   "data": {
		"quizId":"fcf7890f-9c72-46d3-931e-34494307be37"
	},
   "type": "%s"
}
`, command.FinishQuizCommandType)

var eventBodyEmptySelectAnswer = fmt.Sprintf(`
{
   "createdAt":"2023-11-17T04:55:24.059Z",
   "data": {
		"quizId":"fcf7890f-9c72-46d3-931e-34494307be37",
		"questionId":"f5b70d7d-3461-4cf7-978d-2b0caf77db1e",
		"answerId": ""
	},
   "type": "%s"
}
`, command.SelectAnswerCommandType)

var eventBodyEmptyStartQuiz = fmt.Sprintf(`
{
   "createdAt":"2023-11-17T04:55:24.059Z",
   "data": {
		"quizId":"",
		"requiredQuestionsAnswered": ["f5b70d7d-3461-4cf7-978d-2b0caf77db1e"]
	},
   "type": "%s"
}
`, command.StartQuizCommandType)

func TestPutParticipantLambda_Returns200(t *testing.T) {
	environmentCreator := local.NewEnvironmentCreator(config.Test)
	defer environmentCreator.Terminate()

	requestBodys := []string{
		startQuizCommand,
		eventBody2,
		finishQuizComand,
	}

	handler := participant.NewPostParticipantCommandHandler
	for _, requestBody := range requestBodys {
		handlerResponse := environmentCreator.ExecuteLambdaHandlerWithPostBody(
			handler,
			requestBody,
		)
		if handlerResponse.StatusCode != 200 {
			t.Fatalf("lambda did not succeed, status code: %v, %v", handlerResponse.StatusCode, handlerResponse.Body)
		}
	}

}

func TestPutParticipantLambda_InvalidQuizCompletion_Returns500(t *testing.T) {
	environmentCreator := local.NewEnvironmentCreator(config.Test)
	defer environmentCreator.Terminate()

	handler := participant.NewPostParticipantCommandHandler

	environmentCreator.ExecuteLambdaHandlerWithPostBody(
		handler,
		startQuizCommand,
	)

	handlerResponse3 := environmentCreator.ExecuteLambdaHandlerWithPostBody(
		handler,
		finishQuizComand,
	)
	if handlerResponse3.StatusCode != 500 {
		t.Fatalf("lambda return code is not 500 although invalid quiz completion case: %v, %v", handlerResponse3.StatusCode, handlerResponse3.Body)
	}

}

func TestPutParticipantLambda_InvalidQuestionSelection_Returns400(t *testing.T) {
	environmentCreator := local.NewEnvironmentCreator(config.Test)
	defer environmentCreator.Terminate()

	participantCommandHandler := participant.NewPostParticipantCommandHandler

	environmentCreator.ExecuteLambdaHandlerWithPostBody(
		participantCommandHandler,
		startQuizCommand,
	)

	handlerResponse2 := environmentCreator.ExecuteLambdaHandlerWithPostBody(
		participantCommandHandler,
		eventBodyEmptySelectAnswer,
	)
	if handlerResponse2.StatusCode != 400 {
		t.Fatalf("lambda return code is not 400 although invalid quiz completion case: %v, %v", handlerResponse2.StatusCode, handlerResponse2.Body)
	}

}

func TestPutParticipantLambda_InvalidStartQuiz_Returns400(t *testing.T) {
	environmentCreator := local.NewEnvironmentCreator(config.Test)
	defer environmentCreator.Terminate()

	handler := participant.NewPostParticipantCommandHandler

	handlerResponse := environmentCreator.ExecuteLambdaHandlerWithPostBody(
		handler,
		eventBodyEmptyStartQuiz,
	)

	if handlerResponse.StatusCode != 400 {
		t.Fatalf("lambda return code is not 400 although invalid start quiz command: %v, %v", handlerResponse.StatusCode, handlerResponse.Body)
	}

}
