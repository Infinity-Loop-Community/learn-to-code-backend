package lambda

import (
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

type NextJsSecretParser struct {
}

func NewNextJsSecretParser() *NextJsSecretParser {
	return &NextJsSecretParser{}
}

const NextJsSessionTokenCookieKey = "next-auth.session-token"

func (parser *NextJsSecretParser) GetJwtTokenFromRequest(request events.APIGatewayProxyRequest) (string, error) {
	cookieHeader, exists := request.Headers["Cookie"]
	if !exists {
		return "", fmt.Errorf("no cookie presents")
	}

	nextJsAuthSessionToken := parser.getNextJsAuthTokenFromCookie(cookieHeader)

	if nextJsAuthSessionToken == "" {
		return "", fmt.Errorf("no cookie key '%s' present", NextJsSessionTokenCookieKey)
	}

	return nextJsAuthSessionToken, nil
}

func (parser *NextJsSecretParser) getNextJsAuthTokenFromCookie(cookieHeader string) string {
	nextJsAuthSessionToken := ""

	cookies := strings.Split(cookieHeader, "; ")
	for _, cookie := range cookies {
		parts := strings.SplitN(cookie, "=", 2)
		if len(parts) == 2 && parts[0] == "next-auth.session-token" {
			nextJsAuthSessionToken = parts[1]
		}
	}
	return nextJsAuthSessionToken
}
