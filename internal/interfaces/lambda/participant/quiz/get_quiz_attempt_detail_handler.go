package quiz

import (
	"context"
	"fmt"
	"learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/service"
	"strconv"

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

func (gh *GetAttemptDetailHandler) HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	serviceRegistry := service.NewServiceRegistry(ctx, gh.cfg)

	userID, err := serviceRegistry.RequestValidator.ValidateRequest(request)
	if err != nil {
		return serviceRegistry.ResponseCreator.CreateClientErrorResponse(err)
	}

	quizId, ok := request.PathParameters["quizId"]
	if !ok {
		return serviceRegistry.ResponseCreator.CreateClientErrorResponse(fmt.Errorf("quizId required"))
	}

	attemptId, ok := request.PathParameters["attemptId"]
	var attemptIdNumber int64
	if !ok {
		attemptIdNumber, err = strconv.ParseInt(attemptId, 10, 0)
		return serviceRegistry.ResponseCreator.CreateClientErrorResponse(fmt.Errorf("attemptId required"))
	}

	quizAttemptDetailProjection, quizOverviewErr := serviceRegistry.ParticipantApplicationService.GetQuizAttemptDetail(userID, quizId, int(attemptIdNumber))
	if quizOverviewErr != nil {
		return serviceRegistry.ResponseCreator.CreateServerErrorResponse(err)
	}

	quizAttemptDetailResponseObject := serviceRegistry.QuizAttemptDetailMapper.EntityToResponseObject(quizAttemptDetailProjection)

	return serviceRegistry.ResponseCreator.CreateSuccessResponse(quizAttemptDetailResponseObject)
}
