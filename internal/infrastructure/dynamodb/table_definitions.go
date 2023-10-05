package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type TableDefinition struct {
	TableName            string
	KeySchemas           []types.KeySchemaElement
	AttributeDefinitions []types.AttributeDefinition
}

func GetAllTableDefinitions() []TableDefinition {
	return []TableDefinition{
		{
			TableName: "dev_events",
			KeySchemas: []types.KeySchemaElement{
				{
					AttributeName: aws.String("aggregate_id"),
					KeyType:       types.KeyTypeHash,
				},
				{
					AttributeName: aws.String("version"),
					KeyType:       types.KeyTypeRange,
				},
			},
			AttributeDefinitions: []types.AttributeDefinition{
				{
					AttributeName: aws.String("aggregate_id"),
					AttributeType: types.ScalarAttributeTypeS,
				},
				{
					AttributeName: aws.String("version"),
					AttributeType: types.ScalarAttributeTypeN,
				},
			},
		},
	}
}
