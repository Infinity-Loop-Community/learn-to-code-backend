package lambda

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

type NextJsSecretParser struct {
}

func NewNextJsSecretParser() *NextJsSecretParser {
	return &NextJsSecretParser{}
}

const JWTHeaderName = "Next-Auth.Session-Token"

func (s NextJsSecretParser) GetJwtTokenFromRequest(request events.APIGatewayProxyRequest) (string, error) {
	jwtToken, ok := request.Headers[JWTHeaderName]

	if !ok {
		return "", fmt.Errorf("no session token in header")
	}
	return jwtToken, nil
}
