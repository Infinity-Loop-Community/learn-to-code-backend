package local

import (
	"learn-to-code/internal/infrastructure/authentication/jwt"
	"learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/go/util/err"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

type RequestCreator struct {
	cfg config.Config
}

func NewRequestCreator(cfg config.Config) *RequestCreator {
	return &RequestCreator{
		cfg: cfg,
	}
}

func (r *RequestCreator) CreateGETRequest(pathParameters map[string]string, participantID string) events.APIGatewayProxyRequest {
	return events.APIGatewayProxyRequest{
		Resource:                        "",
		Path:                            "",
		HTTPMethod:                      "GET",
		Headers:                         map[string]string{"Cookie": r.createSessionTokenCookie(r.cfg, participantID)},
		MultiValueHeaders:               nil,
		QueryStringParameters:           nil,
		MultiValueQueryStringParameters: nil,
		PathParameters:                  pathParameters,
		StageVariables:                  nil,
		RequestContext:                  events.APIGatewayProxyRequestContext{},
		Body:                            "",
		IsBase64Encoded:                 false,
	}
}

func (r *RequestCreator) CreatePOSTRequest(body string, pathParameters map[string]string, participantID string) events.APIGatewayProxyRequest {
	return events.APIGatewayProxyRequest{
		Resource:                        "",
		Path:                            "",
		HTTPMethod:                      "POST",
		Headers:                         map[string]string{"Cookie": r.createSessionTokenCookie(r.cfg, participantID)},
		MultiValueHeaders:               nil,
		QueryStringParameters:           nil,
		MultiValueQueryStringParameters: nil,
		PathParameters:                  pathParameters,
		StageVariables:                  nil,
		RequestContext:                  events.APIGatewayProxyRequestContext{},
		Body:                            body,
		IsBase64Encoded:                 false,
	}
}

func (r *RequestCreator) createSessionTokenCookie(cfg config.Config, participantID string) string {
	validJwtToken := err.PanicIfError1(jwt.NewValidator(cfg.JwtSecret).CreateToken(participantID))

	cookie := http.Cookie{
		Name:     "next-auth.session-token",
		Value:    validJwtToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour),
	}

	return cookie.String()
}
