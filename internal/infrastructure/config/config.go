package config

import (
	"fmt"
	"os"
)

type Config struct {
	Environment      Environment
	DefaultAwsRegion string
	JwtSecret        string
	CorsAllowOrigin  string
}

const EnvEnvironmentKey = "ENVIRONMENT"
const EnvJwtSecretKey = "JWT_SECRET"
const EnvCorsAllowOriginKey = "CORS_ALLOW_ORIGIN_URL"

func NewConfig() (Config, error) {

	environment, err := ParseEnvironment(os.Getenv(EnvEnvironmentKey))
	if err != nil {
		return Config{}, fmt.Errorf("clould not parse the env variable '%s': %w", EnvEnvironmentKey, err)
	}

	jwtSecret := os.Getenv(EnvJwtSecretKey)
	if jwtSecret == "" {
		return Config{}, fmt.Errorf("missing environment variable '%s'", EnvJwtSecretKey)
	}

	allowOrigin := os.Getenv(EnvCorsAllowOriginKey)
	if allowOrigin == "" {
		return Config{}, fmt.Errorf("missing environment variable '%s'", EnvCorsAllowOriginKey)
	}

	return Config{
		Environment:      environment,
		DefaultAwsRegion: "eu-central-1",
		JwtSecret:        jwtSecret,
		CorsAllowOrigin:  allowOrigin,
	}, err
}
