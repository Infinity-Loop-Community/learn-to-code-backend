package lambda

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func NewResponseCreator(allowOrigin string) *ResponseCreator {
	return &ResponseCreator{
		allowOrigin: allowOrigin,
	}
}

type ResponseCreator struct {
	allowOrigin string
}

func (r *ResponseCreator) CreateSuccessResponse(responeObject any) (events.APIGatewayProxyResponse, error) {
	jsonResponseObject, err := json.Marshal(responeObject)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500,
			Body:    fmt.Sprintf(`{"error": "%s"}`, err),
			Headers: r.getHeaders(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       string(jsonResponseObject),
		Headers:    r.getHeaders(),
		StatusCode: 200,
	}, nil
}

func (r *ResponseCreator) CreateServerErrorResponse(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       fmt.Sprintf(`{"error": "%s"}`, err),
		Headers:    r.getHeaders(),
	}, nil
}

func (r *ResponseCreator) CreateClientErrorResponse(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 400,
		Body:       fmt.Sprintf(`{"error": "%s"}`, err),
		Headers:    r.getHeaders(),
	}, nil
}

func (r *ResponseCreator) CreateNotFoundResponse() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 404,
		Headers:    r.getHeaders(),
	}, nil
}

func (r *ResponseCreator) getHeaders() map[string]string {
	return map[string]string{
		"Access-Control-Allow-Origin":      r.allowOrigin, // Specify your origin here. Use specific domain instead of '*' for production
		"Access-Control-Allow-Methods":     "OPTIONS,GET,PUT,POST,DELETE",
		"Access-Control-Allow-Headers":     "Cookie,Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
		"Access-Control-Allow-Credentials": "true",
	}
}
