package lambda

import (
	"encoding/json"
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

func (r RequestValidator) ValidateRequest(request events.APIGatewayProxyRequest) (Body, string, error) {
	jwtToken, err := r.nextJsSecretParser.GetJwtTokenFromRequest(request)
	if err != nil {
		return Body{}, "", err
	}

	userID, err := r.jwtTokenValidator.ValidateAndGetUserID(jwtToken)
	if err != nil {
		return Body{}, "", err
	}

	body, err := r.getBody(request)
	if err != nil {
		return body, "", err
	}

	return body, userID, nil
}

func (r RequestValidator) getBody(request events.APIGatewayProxyRequest) (Body, error) {
	var body = Body{}

	if request.Body != "" {
		err := json.Unmarshal([]byte(request.Body), &body)
		if err != nil {
			return Body{}, err
		}
	}

	return body, nil
}
