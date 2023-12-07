package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
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

func (validator Validator) CreateToken(subject string) (string, error) {
	claims := &CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "myAppIssuer",
			Subject:   subject,
			Audience:  []string{"myAppClient"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        "uniqueTokenID123",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(validator.secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
