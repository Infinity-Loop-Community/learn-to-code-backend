package db

import (
	"context"
	"fmt"
	dynamodbInfra "learn-to-code/internal/infrastructure/dynamodb"
	errUtils "learn-to-code/internal/infrastructure/go/util/err"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	dynamodbsdk "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

const (
	awsDefaultRegion                     = "eu-region-1"
	dynamoDbPort                         = "8000"
	dynamoDbDefaultProvisionedThroughput = 5
	dynamoDBLocalRepo                    = "amazon/dynamodb-local"
	dynamoDBLocalTag                     = "2.0.0"
	hardShutdownConainterAfterSeconds    = 120
)

func StartDynamoDB() (*dynamodbsdk.Client, func()) {
	pool := errUtils.PanicIfError1(dockertest.NewPool(""))
	container := startDockerContainer(pool)
	dynamoClient := getDynamoClient(pool, container)

	createTables(dynamoClient)

	return dynamoClient, func() {
		if err := pool.Purge(container); err != nil {
			panic(fmt.Errorf("Could not purge DynamoDB: %s", err))
		}
	}
}

func getDynamoClient(pool *dockertest.Pool, container *dockertest.Resource) *dynamodbsdk.Client {
	var dynamoClient *dynamodbsdk.Client

	if err := pool.Retry(func() error {
		cfg, err := config.LoadDefaultConfig(context.Background(),
			config.WithRegion(awsDefaultRegion),
			config.WithEndpointResolverWithOptions(
				aws.EndpointResolverWithOptionsFunc(
					func(service, region string, options ...interface{}) (aws.Endpoint, error) {
						return aws.Endpoint{URL: "http://" + container.GetHostPort(dynamoDbPort+"/tcp")}, nil
					})),
			config.WithCredentialsProvider(
				credentials.StaticCredentialsProvider{
					Value: aws.Credentials{
						AccessKeyID: "dummy", SecretAccessKey: "dummy", SessionToken: "dummy",
						Source: "Hard-coded credentials; values are irrelevant for local DynamoDB",
					},
				}),
		)
		if err != nil {
			return err
		}

		dynamoClient = dynamodbsdk.NewFromConfig(cfg)
		return nil
	}); err != nil {
		panic(fmt.Errorf("Could not connect to the Docker instance of DynamoDB Local: %s", err))
	}

	return dynamoClient
}

func startDockerContainer(pool *dockertest.Pool) *dockertest.Resource {
	runOpt := &dockertest.RunOptions{
		Repository: dynamoDBLocalRepo,
		Tag:        dynamoDBLocalTag,

		PortBindings: map[docker.Port][]docker.PortBinding{
			"0/tcp": {{HostIP: "localhost", HostPort: "8000/tcp"}},
		},
	}
	resource, err := pool.RunWithOptions(runOpt, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})

	errUtils.PanicIfError(resource.Expire(hardShutdownConainterAfterSeconds))

	if err != nil {
		panic(fmt.Errorf("Could not start DynamoDB Local: %s", err))
	}
	println(fmt.Sprintf("Using host:port of '%s'", resource.GetHostPort(dynamoDbPort+"/tcp")))

	return resource
}

func createTables(dynamoDbClient *dynamodbsdk.Client) {

	definitions := dynamodbInfra.GetAllTableDefinitions()

	for _, definition := range definitions {
		exists := errUtils.PanicIfError1(tableExists(definition.TableName, dynamoDbClient))

		if exists {
			errUtils.PanicIfError(deleteTable(definition.TableName, dynamoDbClient))
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

		opt := func(o *dynamodbsdk.Options) { o.RetryMaxAttempts = 10 }

		_ = errUtils.PanicIfError1(dynamoDbClient.CreateTable(context.TODO(), createTableInput, opt))
	}

}

func tableExists(tableName string, dynamoDbClient *dynamodbsdk.Client) (bool, error) {
	_, err := dynamoDbClient.DescribeTable(context.Background(), &dynamodbsdk.DescribeTableInput{
		TableName: aws.String(tableName),
	})

	if err != nil {
		return false, nil
	}

	return true, nil
}

func deleteTable(tableName string, client *dynamodbsdk.Client) error {
	_, err := client.DeleteTable(context.Background(), &dynamodbsdk.DeleteTableInput{
		TableName: aws.String(tableName),
	})

	return err
}
