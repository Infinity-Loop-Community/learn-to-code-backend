package participant

import (
	"context"
	"fmt"
	command "learn-to-code/internal/domain/command"
	"learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/lambda"
	"learn-to-code/internal/infrastructure/service"
	"learn-to-code/internal/interfaces/lambda/course/requestobject"

	"github.com/aws/aws-lambda-go/events"
)

type LambdaHandler struct {
	cfg config.Config
}

func NewPostParticipantCommandHandler(cfg config.Config) *LambdaHandler {
	return &LambdaHandler{cfg: cfg}
}

func (l LambdaHandler) HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	serviceRegistry := service.NewServiceRegistry(ctx, l.cfg)

	userID, err := serviceRegistry.RequestValidator.ValidateRequest(request)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: fmt.Sprintf(`{"error": "%s"}`, err)}, nil
	}

	commandRequest, err := lambda.ReadBody(request.Body, requestobject.Command{})
	if err != nil {
		return serviceRegistry.ResponseCreator.CreateClientErrorResponse(err)
	}

	commandDomainObject := l.mapRequestToCommand(commandRequest)

	err = serviceRegistry.ParticipantApplicationService.ProcessCommand(commandDomainObject, userID)
	if err != nil {
		return serviceRegistry.ResponseCreator.CreateServerErrorResponse(err)
	}

	return serviceRegistry.ResponseCreator.CreateSuccessResponse(commandDomainObject)
}

func (l LambdaHandler) mapRequestToCommand(commandRequest requestobject.Command) command.Command {
	c := command.Command{
		CreatedAt: commandRequest.CreatedAt,
		Data:      commandRequest.Data,
		Type:      commandRequest.Type,
	}
	return c
}
