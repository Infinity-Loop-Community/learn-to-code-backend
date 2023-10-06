package config

import "os"

type Config struct {
	Environment      Environment
	DefaultAwsRegion string
}

func NewConfig() (Config, error) {
	environment, err := ParseEnvironment(os.Getenv("ENVIRONMENT"))
	return Config{
		Environment:      environment,
		DefaultAwsRegion: "eu-central-1",
	}, err
}
