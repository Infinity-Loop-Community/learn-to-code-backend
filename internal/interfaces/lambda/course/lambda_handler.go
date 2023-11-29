package course

import (
	"context"
	"errors"
	"fmt"
	"learn-to-code/internal/domain/quiz/course"
	"learn-to-code/internal/infrastructure/config"
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
