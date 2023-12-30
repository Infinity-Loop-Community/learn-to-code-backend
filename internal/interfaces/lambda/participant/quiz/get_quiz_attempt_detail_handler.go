package quiz

import (
	"context"
	"errors"
	"fmt"
	"learn-to-code/internal/domain/quiz/participant/projection/quizattemptdetail"
	"learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/service"
	"learn-to-code/internal/interfaces/lambda"

	"github.com/aws/aws-lambda-go/events"
)

type GetAttemptDetailHandler struct {
	lambda.HandlerBase
}

func NewGetParticipantQuizAttemptDetailHandler(cfg config.Config, registryOverride service.RegistryOverride) lambda.Handler {
	return &GetAttemptDetailHandler{
		lambda.HandlerBase{
			Cfg:               cfg,
			RegistryOverrides: []service.RegistryOverride{registryOverride},
		},
	}
}

func (gh *GetAttemptDetailHandler) HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	serviceRegistry := service.NewServiceRegistry(ctx, gh.Cfg, gh.RegistryOverrides...)

	userID, err := serviceRegistry.RequestValidator.ValidateRequest(request)
	if err != nil {
		return serviceRegistry.ResponseCreator.CreateClientErrorResponse(err)
	}

	quizID, ok := request.PathParameters["quizId"]
	if !ok {
		return serviceRegistry.ResponseCreator.CreateClientErrorResponse(fmt.Errorf("quizId required"))
	}

	attemptID, ok := request.PathParameters["attemptId"]
	if !ok {
		return serviceRegistry.ResponseCreator.CreateClientErrorResponse(fmt.Errorf("attemptId required"))
	}

	quizAttemptDetailProjection, quizOverviewErr := serviceRegistry.ParticipantApplicationService.GetQuizAttemptDetail(userID, quizID, attemptID)

	if errors.As(quizOverviewErr, &quizattemptdetail.AttemptNotFoundError{}) {
		return serviceRegistry.ResponseCreator.CreateNotFoundResponse()
	} else if quizOverviewErr != nil {
		return serviceRegistry.ResponseCreator.CreateServerErrorResponse(err)
	}

	quizAttemptDetailResponseObject := serviceRegistry.QuizAttemptDetailMapper.EntityToResponseObject(quizAttemptDetailProjection)

	return serviceRegistry.ResponseCreator.CreateSuccessResponse(quizAttemptDetailResponseObject)
}
