package local

import (
	"context"
	"learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/go/util/err"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

type EnvironmentCreator struct {
	Cfg            config.Config
	requestCreator *RequestCreator
}

func NewEnvironmentCreator() *EnvironmentCreator {
	cfg := setupExecutionEnvironment()
	return &EnvironmentCreator{
		Cfg:            cfg,
		requestCreator: NewRequestCreator(cfg),
	}
}

func setupExecutionEnvironment() config.Config {
	os.Setenv(config.EnvEnvironmentKey, string(config.Dev))
	os.Setenv(config.EnvJwtSecretKey, "test")
	os.Setenv(config.EnvCorsAllowOriginKey, "http://localhost:3000")
	cfg := err.PanicIfError1(config.NewConfig())
	return cfg
}

func (ec *EnvironmentCreator) ExecuteLambdaHandler(handleRequest func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)) events.APIGatewayProxyResponse {
	request := err.PanicIfError1(handleRequest(context.Background(), ec.requestCreator.CreateRequest()))
	return request
}
