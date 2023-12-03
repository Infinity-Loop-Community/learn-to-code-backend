package quiz

import (
	"context"
	"learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/service"

	"github.com/aws/aws-lambda-go/events"
)

type GetHandler struct {
	cfg config.Config
}

func NewGetParticipantQuizOverviewHandler(cfg config.Config) *GetHandler {
	return &GetHandler{
		cfg: cfg,
	}
}

func (gh *GetHandler) HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	serviceRegistry := service.NewServiceRegistry(ctx, gh.cfg)

	userID, err := serviceRegistry.RequestValidator.ValidateRequest(request)
	if err != nil {
		return serviceRegistry.ResponseCreator.CreateClientErrorResponse(err)
	}

	quizOverviewProjection, quizOverviewErr := serviceRegistry.ParticipantApplicationService.GetQuizzes(userID)
	if quizOverviewErr != nil {
		return serviceRegistry.ResponseCreator.CreateServerErrorResponse(err)
	}

	return serviceRegistry.ResponseCreator.CreateSuccessResponse(serviceRegistry.QuizOverviewMapper.EntityToResponseObject(quizOverviewProjection))
}
