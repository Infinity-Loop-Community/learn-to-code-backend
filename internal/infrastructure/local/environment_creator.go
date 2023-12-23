package local

import (
	"context"
	"learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/go/util/err"
	"learn-to-code/internal/infrastructure/inmemory"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

type EnvironmentCreator struct {
	Cfg            config.Config
	requestCreator *RequestCreator
}

func NewEnvironmentCreator(environment config.Environment) *EnvironmentCreator {
	cfg := setupExecutionEnvironment(environment)

	return &EnvironmentCreator{
		Cfg:            cfg,
		requestCreator: NewRequestCreator(cfg),
	}
}

func setupExecutionEnvironment(environment config.Environment) config.Config {
	os.Setenv(config.EnvEnvironmentKey, string(environment))
	os.Setenv(config.EnvJwtSecretKey, "test")
	os.Setenv(config.EnvCorsAllowOriginKey, "http://localhost:3000")
	cfg := err.PanicIfError1(config.NewConfig())
	return cfg
}

func (ec *EnvironmentCreator) ExecuteLambdaHandler(handleRequest func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)) events.APIGatewayProxyResponse {
	request := err.PanicIfError1(handleRequest(context.Background(), ec.requestCreator.CreateGETRequest(
		map[string]string{
			"courseId": inmemory.CourseID,
		}, "user123")))
	return request
}

func (ec *EnvironmentCreator) ExecuteLambdaHandlerGETWithPathParametersForUser(participantID string, handleRequest func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error), pathParameters map[string]string) events.APIGatewayProxyResponse {
	request := err.PanicIfError1(handleRequest(context.Background(), ec.requestCreator.CreateGETRequest(pathParameters, participantID)))
	return request
}

func (ec *EnvironmentCreator) ExecuteLambdaHandlerGETWithPathParameters(handleRequest func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error), pathParameters map[string]string) events.APIGatewayProxyResponse {
	request := err.PanicIfError1(handleRequest(context.Background(), ec.requestCreator.CreateGETRequest(pathParameters, "user123")))
	return request
}

func (ec *EnvironmentCreator) ExecuteLambdaHandlerWithPostBodyForUser(
	participantID string,
	handleRequest func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error),
	body string) events.APIGatewayProxyResponse {

	request := err.PanicIfError1(handleRequest(context.Background(), ec.requestCreator.CreatePOSTRequest(body, map[string]string{
		"userId": participantID,
	}, participantID)))
	return request
}

func (ec *EnvironmentCreator) ExecuteLambdaHandlerWithPostBody(
	handleRequest func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error),
	body string) events.APIGatewayProxyResponse {

	request := err.PanicIfError1(handleRequest(context.Background(), ec.requestCreator.CreatePOSTRequest(body, map[string]string{
		"userId": "user123",
	}, "user123")))
	return request
}
