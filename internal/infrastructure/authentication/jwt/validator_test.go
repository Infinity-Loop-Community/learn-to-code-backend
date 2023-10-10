package jwt_test

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	authJwt "learn-to-code/internal/infrastructure/authentication/jwt"
	"testing"
	"time"
)

var sampleSecret = "SecretYouShouldHide"
var sampleSecretKey = []byte(sampleSecret)

func TestValidator_Validate(t *testing.T) {

	t.Run("returns error for invalid secret", func(t *testing.T) {
		validator := authJwt.NewValidator("wrong secret")
		futureTime := time.Now().Add(10 * time.Minute)

		_, err := validator.ValidateAndGetUserId(createRequestWithGeneratedJWT(futureTime))
		assert.ErrorContains(t, err, "signature is invalid")
	})

	t.Run("returns no error for for valid secret and not expired token", func(t *testing.T) {
		validator := authJwt.NewValidator(sampleSecret)
		futureTime := time.Now().Add(10 * time.Minute)

		userId, err := validator.ValidateAndGetUserId(createRequestWithGeneratedJWT(futureTime))
		assert.Nil(t, err)
		assert.Equal(t, "user", userId)
	})

	t.Run("returns error if expired", func(t *testing.T) {
		validator := authJwt.NewValidator(sampleSecret)
		pastTime := time.Now().Add(-10 * time.Minute)

		_, err := validator.ValidateAndGetUserId(createRequestWithGeneratedJWT(pastTime))
		assert.ErrorContains(t, err, "token is expired")
	})
}

func createRequestWithGeneratedJWT(expiresAt time.Time) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, authJwt.CustomClaims{
		Name: "name",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: expiresAt,
			},
			Subject: "user",
		},
	})
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		panic(fmt.Sprintf("Signing Error: %s", err))
	}

	return tokenString

}
