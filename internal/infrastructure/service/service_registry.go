package service

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	dynamodbsdk "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"hello-world/internal/application"
	config2 "hello-world/internal/infrastructure/config"
	"hello-world/internal/infrastructure/dynamodb"
	"hello-world/internal/infrastructure/go/util/err"
)

type Registry struct {
	QuizApplicationService *application.QuizApplicationService
}

func NewServiceRegistry(ctx context.Context, cfg config2.Config) *Registry {
	dynamoDbClient := createDynamoDbClient(ctx, cfg.DefaultAwsRegion)
	participantRepository := dynamodb.NewDynamoDbParticipantRepository(ctx, cfg.Environment, dynamoDbClient)
	quizApplicationService := application.NewQuizApplicationService(participantRepository)

	return &Registry{
		QuizApplicationService: quizApplicationService,
	}
}

func createDynamoDbClient(ctx context.Context, defaultAwsRegion string) *dynamodbsdk.Client {
	dynamoDbConfig := err.PanicIfError1(config.LoadDefaultConfig(ctx))
	dynamoDbConfig.Region = defaultAwsRegion

	return dynamodbsdk.NewFromConfig(dynamoDbConfig)
}
