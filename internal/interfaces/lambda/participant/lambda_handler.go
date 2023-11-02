package participant

import (
	"context"
	"fmt"
	"learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/go/util/uuid"
	"learn-to-code/internal/infrastructure/service"

	"github.com/aws/aws-lambda-go/events"
)

type LambdaHandler struct {
	cfg config.Config
}

func NewLambdaHandler(cfg config.Config) LambdaHandler {
	return LambdaHandler{cfg: cfg}
}

func (l LambdaHandler) HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	serviceRegistry := service.NewServiceRegistry(ctx, l.cfg)

	_, userID, err := serviceRegistry.RequestValidator.ValidateRequest(request)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: fmt.Sprintf(`{"error": "%s"}`, err)}, nil
	}

	err = serviceRegistry.ParticipantApplicationService.StartQuiz(userID, uuid.MustNewRandomAsString())
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: fmt.Sprintf(`{"error": "%s"}`, err)}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       request.Body,
		StatusCode: 200,
	}, nil
}
