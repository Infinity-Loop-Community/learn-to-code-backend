package local

import (
	"context"
	"learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/go/util/err"
	"learn-to-code/internal/infrastructure/inmemory"
	"learn-to-code/internal/infrastructure/service"
	"learn-to-code/internal/infrastructure/testing/db"
	"learn-to-code/internal/interfaces/lambda"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

type EnvironmentCreator struct {
	Cfg              config.Config
	requestCreator   *RequestCreator
	RegistryOverride service.RegistryOverride
	onTerminationFns []func()
}

func NewEnvironmentCreator(environment config.Environment) *EnvironmentCreator {

	var finalOnTerminationFns []func()
	var registryOverride service.RegistryOverride

	if environment == config.Test {
		dynamoDBClient, clean := db.StartDynamoDB()
		finalOnTerminationFns = append(finalOnTerminationFns, clean)

		registryOverride = service.RegistryOverride{DynamoDBClient: dynamoDBClient}
	}

	cfg := setupExecutionEnvironment(environment)

	return &EnvironmentCreator{
		Cfg:              cfg,
		requestCreator:   NewRequestCreator(cfg),
		RegistryOverride: registryOverride,
		onTerminationFns: finalOnTerminationFns,
	}
}

func setupExecutionEnvironment(environment config.Environment) config.Config {
	os.Setenv(config.EnvEnvironmentKey, string(environment))
	os.Setenv(config.EnvJwtSecretKey, "test")
	os.Setenv(config.EnvCorsAllowOriginKey, "http://localhost:3000")
	cfg := err.PanicIfError1(config.NewConfig())
	return cfg
}

func (ec *EnvironmentCreator) Terminate() {
	for _, fn := range ec.onTerminationFns {
		fn()
	}
}

func (ec *EnvironmentCreator) ExecuteLambdaHandler(
	newLambdaHandlerFn func(cfg config.Config, registryOverride service.RegistryOverride) lambda.Handler,
) events.APIGatewayProxyResponse {

	handler := newLambdaHandlerFn(ec.Cfg, ec.RegistryOverride)
	request := err.PanicIfError1(handler.HandleRequest(context.Background(), ec.requestCreator.CreateGETRequest(
		map[string]string{
			"courseId": inmemory.CourseIDFrontendDevelopment,
		}, "user123")))
	return request
}

func (ec *EnvironmentCreator) ExecuteLambdaHandlerGETWithPathParametersForUser(
	participantID string,
	newLambdaHandlerFn func(cfg config.Config, registryOverride service.RegistryOverride) lambda.Handler,
	pathParameters map[string]string,
) events.APIGatewayProxyResponse {

	handler := newLambdaHandlerFn(ec.Cfg, ec.RegistryOverride)
	request := err.PanicIfError1(handler.HandleRequest(context.Background(), ec.requestCreator.CreateGETRequest(pathParameters, participantID)))
	return request
}

func (ec *EnvironmentCreator) ExecuteLambdaHandlerGETWithPathParameters(
	newLambdaHandlerFn func(cfg config.Config, registryOverride service.RegistryOverride) lambda.Handler,
	pathParameters map[string]string,
) events.APIGatewayProxyResponse {

	handler := newLambdaHandlerFn(ec.Cfg, ec.RegistryOverride)
	request := err.PanicIfError1(handler.HandleRequest(context.Background(), ec.requestCreator.CreateGETRequest(pathParameters, "user123")))
	return request
}

func (ec *EnvironmentCreator) ExecuteLambdaHandlerWithPostBodyForUser(
	participantID string,
	newLambdaHandlerFn func(cfg config.Config, registryOverride service.RegistryOverride) lambda.Handler,
	body string,
) events.APIGatewayProxyResponse {

	handler := newLambdaHandlerFn(ec.Cfg, ec.RegistryOverride)
	request := err.PanicIfError1(handler.HandleRequest(context.Background(), ec.requestCreator.CreatePOSTRequest(body, map[string]string{
		"userId": participantID,
	}, participantID)))

	return request
}

func (ec *EnvironmentCreator) ExecuteLambdaHandlerWithPostBody(
	newLambdaHandlerFn func(cfg config.Config, registryOverride service.RegistryOverride) lambda.Handler,
	body string) events.APIGatewayProxyResponse {

	handler := newLambdaHandlerFn(ec.Cfg, ec.RegistryOverride)
	request := err.PanicIfError1(handler.HandleRequest(context.Background(), ec.requestCreator.CreatePOSTRequest(body, map[string]string{
		"userId": "user123",
	}, "user123")))

	return request
}
