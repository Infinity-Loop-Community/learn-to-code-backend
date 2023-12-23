package db

import (
	"context"
	"fmt"
	"learn-to-code/internal/infrastructure/dynamodb"
	errUtils "learn-to-code/internal/infrastructure/go/util/err"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
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

var globalTestcontainer testcontainers.Container

type DynamoStarter struct {
	ctx context.Context
}

func NewDynamoStarter() *DynamoStarter {
	return &DynamoStarter{
		ctx: context.Background(),
	}
}

func (s *DynamoStarter) startContainer() {
	req := testcontainers.ContainerRequest{
		Image:        containerImage,
		Cmd:          []string{"-jar", "DynamoDBLocal.jar", "-inMemory"},
		ExposedPorts: []string{fmt.Sprintf("%s/tcp", dynamoDbPort)},
		WaitingFor:   wait.NewHostPortStrategy(dynamoDbPort),
		Name:         "testcontainer-l2c",
	}

	globalTestcontainer = errUtils.PanicIfError1(testcontainers.GenericContainer(s.ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
		Reuse:            true,
	}))
}

var mutex sync.Mutex

func (s *DynamoStarter) CreateDynamoDbClient(startContainerIfNecessary bool) *dynamodbsdk.Client {

	if startContainerIfNecessary {
		s.startContainerIfNecessary()
	}

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

func (s *DynamoStarter) startContainerIfNecessary() {
	mutex.Lock()
	if globalTestcontainer == nil {
		s.startContainer()
		dynamoDbClient := s.CreateDynamoDbClient(false)
		s.createTables(dynamoDbClient)
	}
	mutex.Unlock()
}

func (s *DynamoStarter) getPort() nat.Port {
	mappedPort := errUtils.PanicIfError1(globalTestcontainer.MappedPort(s.ctx, dynamoDbPort))
	return mappedPort
}

func (s *DynamoStarter) getHost() string {
	return errUtils.PanicIfError1(globalTestcontainer.Host(s.ctx))
}

func (s *DynamoStarter) createTables(dynamoDbClient *dynamodbsdk.Client) {

	definitions := dynamodb.GetAllTableDefinitions()

	for _, definition := range definitions {
		exists := errUtils.PanicIfError1(s.tableExists(definition.TableName, dynamoDbClient))

		if exists {
			errUtils.PanicIfError(s.deleteTable(definition.TableName, dynamoDbClient))
		}

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

func (s *DynamoStarter) tableExists(tableName string, dynamoDbClient *dynamodbsdk.Client) (bool, error) {
	_, err := dynamoDbClient.DescribeTable(context.Background(), &dynamodbsdk.DescribeTableInput{
		TableName: aws.String(tableName),
	})

	if err != nil {
		return false, nil
	}

	return true, nil
}

func (s *DynamoStarter) Terminate() {
	errUtils.PanicIfError(globalTestcontainer.Terminate(s.ctx))
}

func (s *DynamoStarter) deleteTable(tableName string, client *dynamodbsdk.Client) error {
	_, err := client.DeleteTable(context.Background(), &dynamodbsdk.DeleteTableInput{
		TableName: aws.String(tableName),
	})

	return err
}
