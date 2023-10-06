package config

import "os"

type Config struct {
	Environment Environment
}

func NewConfig() (Config, error) {
	environment, err := ParseEnvironment(os.Getenv("ENVIRONMENT"))
	return Config{Environment: environment}, err
}
