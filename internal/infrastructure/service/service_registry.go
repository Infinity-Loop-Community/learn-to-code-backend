package service

import (
	"context"
	"learn-to-code/internal/application"
	"learn-to-code/internal/domain/command"
	authJwt "learn-to-code/internal/infrastructure/authentication/jwt"
	config2 "learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/dynamodb"
	"learn-to-code/internal/infrastructure/go/util/err"
	"learn-to-code/internal/infrastructure/inmemory"
	"learn-to-code/internal/infrastructure/lambda"
	"learn-to-code/internal/interfaces/lambda/course/mapper"

	"github.com/aws/aws-sdk-go-v2/config"
	dynamodbsdk "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Registry struct {
	ParticipantApplicationService *application.ParticipantApplicationService
	CourseApplicationService      *application.CourseApplicationService
	RequestValidator              *lambda.RequestValidator
	CourseMapper                  *mapper.CourseMapper
	ResponseCreator               *lambda.ResponseCreator
}

func NewServiceRegistry(ctx context.Context, cfg config2.Config) *Registry {
	dynamoDbClient := createDynamoDbClient(ctx, cfg.DefaultAwsRegion)

	nextJsSecretParser := lambda.NewNextJsSecretParser()
	jwtTokenValidator := authJwt.NewValidator(cfg.JwtSecret)
	requestValidator := lambda.NewRequestValidator(nextJsSecretParser, jwtTokenValidator)
	responseCreator := lambda.NewResponseCreator(cfg.CorsAllowOrigin)

	startQuizToEventMapper := command.NewParticipantCommandApplier()

	participantRepositoryFactory := dynamodb.NewParticipantRepositoryFactory(cfg.Environment, dynamoDbClient)
	participantRepository := participantRepositoryFactory.NewRepository(ctx)
	participantApplicationService := application.NewPartcipantApplicationService(participantRepository, startQuizToEventMapper)

	courseRepository := inmemory.NewCourseRepository()
	courseApplicationService := application.NewCourseApplicationService(courseRepository)
	courseMapper := mapper.NewCourseMapper()

	return &Registry{
		ParticipantApplicationService: participantApplicationService,
		CourseApplicationService:      courseApplicationService,
		CourseMapper:                  courseMapper,
		RequestValidator:              requestValidator,
		ResponseCreator:               responseCreator,
	}
}

func createDynamoDbClient(ctx context.Context, defaultAwsRegion string) *dynamodbsdk.Client {
	dynamoDbConfig := err.PanicIfError1(config.LoadDefaultConfig(ctx))
	dynamoDbConfig.Region = defaultAwsRegion

	return dynamodbsdk.NewFromConfig(dynamoDbConfig)
}
