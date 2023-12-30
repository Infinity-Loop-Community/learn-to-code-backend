package quiz

import (
	"context"
	"learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/service"
	"learn-to-code/internal/interfaces/lambda"

	"github.com/aws/aws-lambda-go/events"
)

type GetOverviewHandler struct {
	lambda.HandlerBase
}

func NewGetParticipantQuizOverviewHandler(cfg config.Config, registryOverride service.RegistryOverride) lambda.Handler {
	return &GetOverviewHandler{
		lambda.HandlerBase{
			Cfg:               cfg,
			RegistryOverrides: []service.RegistryOverride{registryOverride},
		},
	}
}

func (gh *GetOverviewHandler) HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	serviceRegistry := service.NewServiceRegistry(ctx, gh.Cfg, gh.RegistryOverrides...)

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
