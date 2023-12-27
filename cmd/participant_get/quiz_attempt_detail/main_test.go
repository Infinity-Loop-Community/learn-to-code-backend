package main

import (
	"fmt"
	"learn-to-code/internal/domain/command"
	"learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/go/util/uuid"
	"learn-to-code/internal/infrastructure/inmemory"
	"learn-to-code/internal/infrastructure/local"
	"learn-to-code/internal/infrastructure/testing/json"
	"learn-to-code/internal/interfaces/lambda/participant"
	"learn-to-code/internal/interfaces/lambda/participant/quiz"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

var startQuizPayload = fmt.Sprintf(`
{
   "createdAt":"2023-11-17T04:55:24.059Z",
   "data": {
		"quizId":"%s",
		"requiredQuestionsAnswered": ["%s"]
	},
   "type": "%s"
}
`, inmemory.QuizID, inmemory.FirstQuestionID, command.StartQuizCommandType)

var selectAnswerPayload = fmt.Sprintf(`
{
   "createdAt":"2023-11-17T04:55:24.059Z",
   "data": {
		"quizId":"%s",
		"questionId":"%s",
		"answerId": "%s"
	},
   "type": "%s"
}
`, inmemory.QuizID, inmemory.FirstQuestionID, inmemory.FirstCorrectAnswerID, command.SelectAnswerCommandType)

var finishQuizPayload = fmt.Sprintf(`
{
   "createdAt":"2023-11-17T04:55:24.059Z",
   "data": {
		"quizId":"%s"
	},
   "type": "%s"
}
`, inmemory.QuizID, command.FinishQuizCommandType)

func TestGetAttemptDetail_Returns200(t *testing.T) {

	getDetailResponse := requestQuizAttemptDetailByAttemptID("1")

	if getDetailResponse.StatusCode != 200 {
		t.Fatalf("lambda did not succeed, status code: %v, body: %s", getDetailResponse.StatusCode, getDetailResponse.Body)
	}

	questionWithAnswer := json.GetJSONPathValue(getDetailResponse, "$.questionsWithAnswer").(map[string]interface{})

	if questionWithAnswer[inmemory.FirstQuestionID] != inmemory.FirstCorrectAnswerID {
		t.Fatalf("question detail response does not contain the provided answer")
	}
}

func TestGetAttemptDetail_ReturnsForLatest(t *testing.T) {

	getDetailResponse := requestQuizAttemptDetailByAttemptID("latest")

	if getDetailResponse.StatusCode != 200 {
		t.Fatalf("lambda did not succeed, status code: %v, body: %s", getDetailResponse.StatusCode, getDetailResponse.Body)
	}
}

func TestGetAttemptDetail_ReturnStateOngoingForRestartedQuiz(t *testing.T) {

	getDetailResponse := requestQuizAttemptDetailWithRestartedQuizByAttemptID("latest")

	attemptStatus := json.GetJSONPathValue(getDetailResponse, "$.attemptStatus").(string)

	if attemptStatus != "ongoing" {
		t.Fatalf("question detail response does not contain an attempt status 'ongoing'")
	}
}

func TestGetAttemptDetail_ReturnsAttemptStatus(t *testing.T) {

	getDetailResponse := requestQuizAttemptDetailByAttemptID("1")

	attemptStatus := json.GetJSONPathValue(getDetailResponse, "$.attemptStatus").(string)

	if attemptStatus != "ongoing" {
		t.Fatalf("question detail response does not contain an attempt status 'ongoing'")
	}
}

func TestGetAttemptDetail_ReturnsID(t *testing.T) {

	getDetailResponse := requestQuizAttemptDetailByAttemptID("1")

	attemptID := json.GetJSONPathValue(getDetailResponse, "$.attemptId").(float64)

	if attemptID != 1 {
		t.Fatalf("question detail response does not contain an ID '1', but '%f' instead", attemptID)
	}
}

func TestGetAttemptDetail_ReturnsAttemptResult(t *testing.T) {

	getDetailResponse := requestQuizAttemptDetailWithFinishedQuizByAttemptID("1")

	pass := json.GetJSONPathValue(getDetailResponse, "$.attemptResult.pass").(bool)

	if pass != true {
		t.Fatalf("expected quiz with only correct answers to pass but did not pass")
	}
}

func TestGetAttemptDetail_UnknownAttempt_Returns404(t *testing.T) {

	getDetailResponse := requestQuizAttemptDetailByAttemptID("0")

	if getDetailResponse.StatusCode != 404 {
		t.Fatalf("lambda did not return 404 for unknown attempt ID, status code: %v, body: %s", getDetailResponse.StatusCode, getDetailResponse.Body)
	}
}

func requestQuizAttemptDetailByAttemptID(attemptID string) events.APIGatewayProxyResponse {
	environmentCreator := local.NewEnvironmentCreator(config.Test)
	defer environmentCreator.Terminate()

	environmentCreator.ExecuteLambdaHandlerWithPostBody(participant.NewPostParticipantCommandHandler(environmentCreator.Cfg).HandleRequest, startQuizPayload)
	environmentCreator.ExecuteLambdaHandlerWithPostBody(participant.NewPostParticipantCommandHandler(environmentCreator.Cfg).HandleRequest, selectAnswerPayload)

	getDetailResponse := environmentCreator.ExecuteLambdaHandlerGETWithPathParameters(
		quiz.NewGetParticipantQuizAttemptDetailHandler(environmentCreator.Cfg).HandleRequest,
		map[string]string{
			"quizId":    inmemory.QuizID,
			"attemptId": attemptID,
		})
	return getDetailResponse
}

func requestQuizAttemptDetailWithFinishedQuizByAttemptID(attemptID string) events.APIGatewayProxyResponse {
	environmentCreator := local.NewEnvironmentCreator(config.Test)
	defer environmentCreator.Terminate()

	participantID := uuid.MustNewRandomAsString()

	environmentCreator.ExecuteLambdaHandlerWithPostBodyForUser(participantID, participant.NewPostParticipantCommandHandler(environmentCreator.Cfg).HandleRequest, startQuizPayload)
	environmentCreator.ExecuteLambdaHandlerWithPostBodyForUser(participantID, participant.NewPostParticipantCommandHandler(environmentCreator.Cfg).HandleRequest, selectAnswerPayload)
	environmentCreator.ExecuteLambdaHandlerWithPostBodyForUser(participantID, participant.NewPostParticipantCommandHandler(environmentCreator.Cfg).HandleRequest, finishQuizPayload)

	getDetailResponse := environmentCreator.ExecuteLambdaHandlerGETWithPathParametersForUser(
		participantID,
		quiz.NewGetParticipantQuizAttemptDetailHandler(environmentCreator.Cfg).HandleRequest,
		map[string]string{
			"quizId":    inmemory.QuizID,
			"attemptId": attemptID,
		})
	return getDetailResponse
}

func requestQuizAttemptDetailWithRestartedQuizByAttemptID(attemptID string) events.APIGatewayProxyResponse {
	environmentCreator := local.NewEnvironmentCreator(config.Test)
	defer environmentCreator.Terminate()

	participantID := uuid.MustNewRandomAsString()

	environmentCreator.ExecuteLambdaHandlerWithPostBodyForUser(participantID, participant.NewPostParticipantCommandHandler(environmentCreator.Cfg).HandleRequest, startQuizPayload)
	environmentCreator.ExecuteLambdaHandlerWithPostBodyForUser(participantID, participant.NewPostParticipantCommandHandler(environmentCreator.Cfg).HandleRequest, selectAnswerPayload)
	environmentCreator.ExecuteLambdaHandlerWithPostBodyForUser(participantID, participant.NewPostParticipantCommandHandler(environmentCreator.Cfg).HandleRequest, finishQuizPayload)
	environmentCreator.ExecuteLambdaHandlerWithPostBodyForUser(participantID, participant.NewPostParticipantCommandHandler(environmentCreator.Cfg).HandleRequest, startQuizPayload)

	getDetailResponse := environmentCreator.ExecuteLambdaHandlerGETWithPathParametersForUser(
		participantID,
		quiz.NewGetParticipantQuizAttemptDetailHandler(environmentCreator.Cfg).HandleRequest,
		map[string]string{
			"quizId":    inmemory.QuizID,
			"attemptId": attemptID,
		})
	return getDetailResponse
}
