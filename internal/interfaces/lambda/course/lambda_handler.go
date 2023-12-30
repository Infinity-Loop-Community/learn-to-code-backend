package course

import (
	"context"
	"errors"
	"fmt"
	"learn-to-code/internal/domain/quiz/course"
	"learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/service"
	"learn-to-code/internal/interfaces/lambda"

	"github.com/aws/aws-lambda-go/events"
)

type LambdaHandler struct {
	lambda.HandlerBase
}

func NewLambdaHandler(cfg config.Config, registryOverride service.RegistryOverride) lambda.Handler {
	return &LambdaHandler{
		lambda.HandlerBase{
			Cfg:               cfg,
			RegistryOverrides: []service.RegistryOverride{registryOverride},
		},
	}
}

func (l *LambdaHandler) HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	serviceRegistry := service.NewServiceRegistry(ctx, l.Cfg, l.RegistryOverrides...)

	_, err := serviceRegistry.RequestValidator.ValidateRequest(request)
	if err != nil {
		return serviceRegistry.ResponseCreator.CreateClientErrorResponse(err)
	}

	courseID, exists := request.PathParameters["courseId"]
	if !exists {
		return serviceRegistry.ResponseCreator.CreateClientErrorResponse(fmt.Errorf("courseId required"))
	}

	resultCourse, err := serviceRegistry.CourseApplicationService.GetCourseByID(courseID)
	if errors.Is(err, course.ErrCourseNotFound) {
		return serviceRegistry.ResponseCreator.CreateNotFoundResponse()
	}
	if err != nil {
		return serviceRegistry.ResponseCreator.CreateClientErrorResponse(err)
	}

	responseObject := serviceRegistry.CourseMapper.EntityToResponseObject(resultCourse)

	return serviceRegistry.ResponseCreator.CreateSuccessResponse(responseObject)
}
