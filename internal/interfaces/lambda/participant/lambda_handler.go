package participant

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"hello-world/internal/infrastructure/config"
	"hello-world/internal/infrastructure/go/util/uuid"
	"hello-world/internal/infrastructure/service"
)

type LambdaHandler struct {
	cfg config.Config
}

func NewLambdaHandler(cfg config.Config) LambdaHandler {
	return LambdaHandler{cfg: cfg}
}

type Body struct {
	Input string `json:"input"`
}

func (l LambdaHandler) HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var body Body

	err := json.Unmarshal([]byte(request.Body), &body)

	serviceRegistry := service.NewServiceRegistry(ctx, l.cfg)
	serviceRegistry.QuizApplicationService.StartQuiz(uuid.MustNewRandomAsString(), uuid.MustNewRandomAsString())

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: `{"msg": "ready body error, Invalid JSON"}`}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       request.Body,
		StatusCode: 200,
	}, nil
}
