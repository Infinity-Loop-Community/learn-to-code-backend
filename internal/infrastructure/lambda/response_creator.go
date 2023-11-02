package lambda

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func NewResponseCreator() *ResponseCreator {
	return &ResponseCreator{}
}

type ResponseCreator struct {
}

func (r *ResponseCreator) CreateSuccessResponse(responeObject any) (events.APIGatewayProxyResponse, error) {
	jsonResponseObject, err := json.Marshal(responeObject)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: fmt.Sprintf(`{"error": "%s"}`, err)}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       string(jsonResponseObject),
		StatusCode: 200,
	}, nil
}

func (r *ResponseCreator) CreateClientErrorResponse(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{StatusCode: 400, Body: fmt.Sprintf(`{"error": "%s"}`, err)}, nil
}

func (r *ResponseCreator) CreateNotFoundResponse() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{StatusCode: 404}, nil
}
