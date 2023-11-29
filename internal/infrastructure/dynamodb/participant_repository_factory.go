package dynamodb

import (
	"context"
	"learn-to-code/internal/domain/quiz/participant"
	"learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/inmemory"

	dynamodbsdk "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type ParticipantRepositoryFactory struct {
	env            config.Environment
	dynamoDBClient *dynamodbsdk.Client
}

func NewParticipantRepositoryFactory(env config.Environment, dynamoDBClient *dynamodbsdk.Client) *ParticipantRepositoryFactory {
	return &ParticipantRepositoryFactory{
		env:            env,
		dynamoDBClient: dynamoDBClient,
	}
}

func (f *ParticipantRepositoryFactory) NewRepository(ctx context.Context) participant.Repository {
	if f.env == config.Test {
		return inmemory.NewParticipantRepository()
	}

	return NewDynamoDbParticipantRepository(ctx, f.env, f.dynamoDBClient)
}
