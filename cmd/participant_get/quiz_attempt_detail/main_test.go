package main

import (
	"fmt"
	"learn-to-code/internal/domain/command/data"
	"learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/inmemory"
	"learn-to-code/internal/infrastructure/local"
	"learn-to-code/internal/infrastructure/testing/json"
	"learn-to-code/internal/interfaces/lambda/participant"
	"learn-to-code/internal/interfaces/lambda/participant/quiz"
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

func TestGetQuizOverview_Returns200(t *testing.T) {

	environmentCreator := local.NewEnvironmentCreator(config.Test)
	environmentCreator.ExecuteLambdaHandlerWithPostBody(participant.NewPostParticipantCommandHandler(environmentCreator.Cfg).HandleRequest, eventBody)
	environmentCreator.ExecuteLambdaHandlerWithPostBody(participant.NewPostParticipantCommandHandler(environmentCreator.Cfg).HandleRequest, eventBody2)

	getDetailResponse := environmentCreator.ExecuteLambdaHandlerGETWithPathParameters(
		quiz.NewGetParticipantQuizAttemptDetailHandler(environmentCreator.Cfg).HandleRequest,
		map[string]string{
			"quizId":    inmemory.QuizID,
			"attemptId": "0",
		})

	if getDetailResponse.StatusCode != 200 {
		t.Fatalf("lambda did not succeed, status code: %v, body: %s", getDetailResponse.StatusCode, getDetailResponse.Body)
	}

	questionWithAnswer := json.GetJSONPathValue(getDetailResponse, "$.questionsWithAnswer").(map[string]interface{})

	if questionWithAnswer["f5b70d7d-3461-4cf7-978d-2b0caf77db1e"] != "1eb5a02f-562a-409d-8822-a240e0886485" {
		t.Fatalf("question detail response does not contain the provided answer")
	}
}
