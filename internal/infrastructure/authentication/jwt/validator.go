package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

type Validator struct {
	secret string
}

func NewValidator(secret string) *Validator {
	return &Validator{
		secret: secret,
	}
}

func (validator Validator) ValidateAndGetUserID(jwtToken string) (string, error) {

	if validator.secret == "" {
		panic("no next auth secret")
	}

	token, err := jwt.ParseWithClaims(jwtToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(validator.secret), nil
	})

	if err != nil {
		return "", err
	}

	subject, err := token.Claims.GetSubject()

	return subject, err
}
