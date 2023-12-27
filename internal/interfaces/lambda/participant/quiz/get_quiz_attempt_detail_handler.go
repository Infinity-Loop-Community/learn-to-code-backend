package quiz

import (
	"context"
	"errors"
	"fmt"
	"learn-to-code/internal/domain/quiz/participant/projection/quizattemptdetail"
	"learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/service"

	"github.com/aws/aws-lambda-go/events"
)

type GetAttemptDetailHandler struct {
	cfg config.Config
}

func NewGetParticipantQuizAttemptDetailHandler(cfg config.Config) *GetAttemptDetailHandler {
	return &GetAttemptDetailHandler{
		cfg: cfg,
	}
}

func (gh *GetAttemptDetailHandler) HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest, registryOverrides ...service.RegistryOverride) (events.APIGatewayProxyResponse, error) {
	serviceRegistry := service.NewServiceRegistry(ctx, gh.cfg, registryOverrides...)

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
