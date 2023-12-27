package local

import (
	"context"
	"learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/go/util/err"
	"learn-to-code/internal/infrastructure/inmemory"
	"learn-to-code/internal/infrastructure/service"
	"learn-to-code/internal/infrastructure/testing/db"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

type EnvironmentCreator struct {
	Cfg               config.Config
	requestCreator    *RequestCreator
	registryOverrides []service.RegistryOverride
	onTerminationFns  []func()
}

func NewEnvironmentCreator(environment config.Environment, registryOverrides ...service.RegistryOverride) *EnvironmentCreator {

	var finalOnTerminationFns []func()
	var finalRegistryOverrides = registryOverrides

	if environment == config.Test && len(registryOverrides) == 0 {
		dynamoDBClient, clean := db.StartDynamoDB()
		finalOnTerminationFns = append(finalOnTerminationFns, clean)

		finalRegistryOverrides = append(finalRegistryOverrides, service.RegistryOverride{DynamoDBClient: dynamoDBClient})
	}

	cfg := setupExecutionEnvironment(environment)

	return &EnvironmentCreator{
		Cfg:               cfg,
		requestCreator:    NewRequestCreator(cfg),
		registryOverrides: finalRegistryOverrides,
		onTerminationFns:  finalOnTerminationFns,
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
	handleRequest func(ctx context.Context, request events.APIGatewayProxyRequest, registryOverrides ...service.RegistryOverride) (events.APIGatewayProxyResponse, error),
) events.APIGatewayProxyResponse {

	request := err.PanicIfError1(handleRequest(context.Background(), ec.requestCreator.CreateGETRequest(
		map[string]string{
			"courseId": inmemory.CourseID,
		}, "user123"), ec.registryOverrides...))
	return request
}

func (ec *EnvironmentCreator) ExecuteLambdaHandlerGETWithPathParametersForUser(
	participantID string,
	handleRequest func(ctx context.Context, request events.APIGatewayProxyRequest, registryOverrides ...service.RegistryOverride) (events.APIGatewayProxyResponse, error),
	pathParameters map[string]string,
) events.APIGatewayProxyResponse {

	request := err.PanicIfError1(handleRequest(context.Background(), ec.requestCreator.CreateGETRequest(pathParameters, participantID), ec.registryOverrides...))
	return request
}

func (ec *EnvironmentCreator) ExecuteLambdaHandlerGETWithPathParameters(
	handleRequest func(ctx context.Context, request events.APIGatewayProxyRequest, registryOverrides ...service.RegistryOverride) (events.APIGatewayProxyResponse, error),
	pathParameters map[string]string,
) events.APIGatewayProxyResponse {

	request := err.PanicIfError1(handleRequest(context.Background(), ec.requestCreator.CreateGETRequest(pathParameters, "user123"), ec.registryOverrides...))
	return request
}

func (ec *EnvironmentCreator) ExecuteLambdaHandlerWithPostBodyForUser(
	participantID string,
	handleRequest func(ctx context.Context, request events.APIGatewayProxyRequest, registryOverrides ...service.RegistryOverride) (events.APIGatewayProxyResponse, error),
	body string,
) events.APIGatewayProxyResponse {

	request := err.PanicIfError1(handleRequest(context.Background(), ec.requestCreator.CreatePOSTRequest(body, map[string]string{
		"userId": participantID,
	}, participantID), ec.registryOverrides...))

	return request
}

func (ec *EnvironmentCreator) ExecuteLambdaHandlerWithPostBody(
	handleRequest func(ctx context.Context, request events.APIGatewayProxyRequest, registryOverrides ...service.RegistryOverride) (events.APIGatewayProxyResponse, error),
	body string) events.APIGatewayProxyResponse {

	request := err.PanicIfError1(handleRequest(context.Background(), ec.requestCreator.CreatePOSTRequest(body, map[string]string{
		"userId": "user123",
	}, "user123"), ec.registryOverrides...))

	return request
}
