package config

import (
	"fmt"
	"os"
)

type Config struct {
	Environment      Environment
	DefaultAwsRegion string
	JwtSecret        string
}

func NewConfig() (Config, error) {
	environment, err := ParseEnvironment(os.Getenv("ENVIRONMENT"))
	if err != nil {
		return Config{}, err
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return Config{}, fmt.Errorf("missing environment variable 'JWT_SECRET'")
	}

	return Config{
		Environment:      environment,
		DefaultAwsRegion: "eu-central-1",
		JwtSecret:        jwtSecret,
	}, err
}
