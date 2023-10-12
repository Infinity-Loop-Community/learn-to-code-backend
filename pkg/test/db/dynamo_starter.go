package db

import (
	"context"
	"fmt"
	"learn-to-code/internal/infrastructure/dynamodb"
	errUtils "learn-to-code/internal/infrastructure/go/util/err"
	"time"

	"github.com/aws/aws-sdk-go-v2/credentials"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	dynamodbsdk "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	awsDefaultRegion                     = "eu-region-1"
	dynamoDbPort                         = "8000"
	dynamoDbDefaultProvisionedThroughput = 5
	containerImage                       = "amazon/dynamodb-local:2.0.0"
)

type DynamoStarter struct {
	ctx context.Context
	tc  testcontainers.Container
}

func NewDynamoStarter() *DynamoStarter {
	return &DynamoStarter{
		ctx: context.Background(),
		tc:  nil,
	}
}

func (s *DynamoStarter) Start() *dynamodbsdk.Client {
	s.startContainer()
	dynamoDbClient := s.createDynamoDbClient()

	fmt.Printf("Sleep 5 seconds to ensure database is up and running.")
	time.Sleep(time.Second * 5)
	fmt.Printf("Sleep done.")

	s.createTables(dynamoDbClient)
	return dynamoDbClient
}

func (s *DynamoStarter) startContainer() {
	req := testcontainers.ContainerRequest{
		Image:        containerImage,
		Cmd:          []string{"-jar", "DynamoDBLocal.jar", "-inMemory"},
		ExposedPorts: []string{fmt.Sprintf("%s/tcp", dynamoDbPort)},
		WaitingFor:   wait.NewHostPortStrategy(dynamoDbPort),
	}

	s.tc = errUtils.PanicIfError1(testcontainers.GenericContainer(s.ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	}))
}

func (s *DynamoStarter) createDynamoDbClient() *dynamodbsdk.Client {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		customEndpoint := fmt.Sprintf("http://%s:%d", s.getHost(), s.getPort().Int())
		endpoint := aws.Endpoint{
			URL:           customEndpoint,
			SigningRegion: awsDefaultRegion,
		}

		return endpoint, nil
	})

	cfg := errUtils.PanicIfError1(config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsDefaultRegion),
		config.WithEndpointResolverWithOptions(customResolver),

		// disable authentication
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("test", "test", "test")),
	))

	return dynamodbsdk.NewFromConfig(cfg)
}

func (s *DynamoStarter) getPort() nat.Port {
	return errUtils.PanicIfError1(s.tc.MappedPort(s.ctx, dynamoDbPort))
}

func (s *DynamoStarter) getHost() string {
	return errUtils.PanicIfError1(s.tc.Host(s.ctx))
}

func (s *DynamoStarter) createTables(dynamoDbClient *dynamodbsdk.Client) {

	definitions := dynamodb.GetAllTableDefinitions()

	for _, definition := range definitions {
		createTableInput := &dynamodbsdk.CreateTableInput{
			TableName:            aws.String(definition.TableName),
			KeySchema:            definition.KeySchemas,
			AttributeDefinitions: definition.AttributeDefinitions,
			ProvisionedThroughput: &types.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(dynamoDbDefaultProvisionedThroughput),
				WriteCapacityUnits: aws.Int64(dynamoDbDefaultProvisionedThroughput),
			},
		}

		_ = errUtils.PanicIfError1(dynamoDbClient.CreateTable(context.TODO(), createTableInput))
	}

}

func (s *DynamoStarter) Terminate() {
	errUtils.PanicIfError(s.tc.Terminate(s.ctx))
}
