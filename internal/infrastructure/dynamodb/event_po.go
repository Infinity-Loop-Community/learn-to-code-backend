package dynamodb

import "time"

type EventPo struct {
	AggregateID string    `dynamodbav:"aggregate_id"`
	Type        string    `dynamodbav:"type"`
	Version     uint      `dynamodbav:"version"`
	Payload     string    `dynamodbav:"payload"`
	CreatedAt   time.Time `dynamodbav:"created_at"`
}
