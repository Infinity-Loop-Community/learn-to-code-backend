package service

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	dynamodbsdk "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"learn-to-code/internal/application"
	authJwt "learn-to-code/internal/infrastructure/authentication/jwt"
	config2 "learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/dynamodb"
	"learn-to-code/internal/infrastructure/go/util/err"
	"learn-to-code/internal/infrastructure/lambda"
)

type Registry struct {
	QuizApplicationService *application.QuizApplicationService
	RequestValidator       *lambda.RequestValidator
}

func NewServiceRegistry(ctx context.Context, cfg config2.Config) *Registry {
	dynamoDbClient := createDynamoDbClient(ctx, cfg.DefaultAwsRegion)

	nextJsSecretParser := lambda.NewNextJsSecretParser()
	jwtTokenValidator := authJwt.NewValidator(cfg.JwtSecret)
	requestValidator := lambda.NewRequestValidator(nextJsSecretParser, jwtTokenValidator)

	participantRepository := dynamodb.NewDynamoDbParticipantRepository(ctx, cfg.Environment, dynamoDbClient)

	quizApplicationService := application.NewQuizApplicationService(participantRepository)

	return &Registry{
		QuizApplicationService: quizApplicationService,
		RequestValidator:       requestValidator,
	}
}

func createDynamoDbClient(ctx context.Context, defaultAwsRegion string) *dynamodbsdk.Client {
	dynamoDbConfig := err.PanicIfError1(config.LoadDefaultConfig(ctx))
	dynamoDbConfig.Region = defaultAwsRegion

	return dynamodbsdk.NewFromConfig(dynamoDbConfig)
}
