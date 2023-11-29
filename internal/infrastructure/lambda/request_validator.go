package lambda

import (
	authJwt "learn-to-code/internal/infrastructure/authentication/jwt"

	"github.com/aws/aws-lambda-go/events"
)

type RequestValidator struct {
	nextJsSecretParser *NextJsSecretParser
	jwtTokenValidator  *authJwt.Validator
}

func NewRequestValidator(nextJsSecretParser *NextJsSecretParser, jwtTokenValidator *authJwt.Validator) *RequestValidator {
	return &RequestValidator{
		nextJsSecretParser: nextJsSecretParser,
		jwtTokenValidator:  jwtTokenValidator,
	}
}

type Body struct {
	Input string `json:"input"`
}

func (r RequestValidator) ValidateRequest(request events.APIGatewayProxyRequest) (string, error) {
	jwtToken, err := r.nextJsSecretParser.GetJwtTokenFromRequest(request)
	if err != nil {
		return "", err
	}

	userID, err := r.jwtTokenValidator.ValidateAndGetUserID(jwtToken)
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	return userID, nil
}
