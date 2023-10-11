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

func (s NextJsSecretParser) GetJwtTokenFromRequest(request events.APIGatewayProxyRequest) (string, error) {
	jwtToken, ok := request.Headers["next-auth.session-token"]
	if !ok {
		return "", fmt.Errorf("no session token in header")
	}
	return jwtToken, nil
}
