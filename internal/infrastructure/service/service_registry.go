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
	mapper2 "learn-to-code/internal/interfaces/lambda/participant/quiz/mapper"

	"github.com/aws/aws-sdk-go-v2/config"
	dynamodbsdk "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Registry struct {
	RequestValidator *lambda.RequestValidator
	ResponseCreator  *lambda.ResponseCreator

	ParticipantApplicationService *application.ParticipantApplicationService
	QuizOverviewMapper            *mapper2.QuizOverviewMapper

	CourseApplicationService *application.CourseApplicationService
	CourseMapper             *mapper.CourseMapper
	QuizAttemptDetailMapper  *mapper2.QuizAttemptDetailMapper
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
	quizOverviewMapper := mapper2.NewQuizOverviewMapper()
	quizAttemptDetailMapper := mapper2.NewQuizAttemptDetailMapper()

	courseRepository := inmemory.NewCourseRepository()
	courseApplicationService := application.NewCourseApplicationService(courseRepository)
	courseMapper := mapper.NewCourseMapper()

	return &Registry{
		ParticipantApplicationService: participantApplicationService,
		CourseApplicationService:      courseApplicationService,
		CourseMapper:                  courseMapper,
		QuizOverviewMapper:            quizOverviewMapper,
		QuizAttemptDetailMapper:       quizAttemptDetailMapper,
		RequestValidator:              requestValidator,
		ResponseCreator:               responseCreator,
	}
}

func createDynamoDbClient(ctx context.Context, defaultAwsRegion string) *dynamodbsdk.Client {
	dynamoDbConfig := err.PanicIfError1(config.LoadDefaultConfig(ctx))
	dynamoDbConfig.Region = defaultAwsRegion

	return dynamodbsdk.NewFromConfig(dynamoDbConfig)
}
