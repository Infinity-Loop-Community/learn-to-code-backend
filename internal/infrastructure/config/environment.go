package config

import (
	"fmt"
)

type Environment string

const (
	Dev  Environment = "Dev"
	Prod Environment = "Prod"
)

func ParseEnvironment(envVar string) (Environment, error) {
	switch envVar {
	case string(Dev):
		return Dev, nil
	case string(Prod):
		return Prod, nil
	default:
		return "", fmt.Errorf("unsupported environment value '%s'", envVar)
	}
}
