package lambda

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type Handler interface {
	HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}
