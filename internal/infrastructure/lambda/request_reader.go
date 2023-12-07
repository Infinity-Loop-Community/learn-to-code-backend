package lambda

import (
	"encoding/json"
)

func ReadBody[T Validatable](body string, commandRequest T) (T, error) {
	err := json.Unmarshal([]byte(body), &commandRequest)
	if err != nil {
		return commandRequest, err
	}

	err = commandRequest.Validate()
	if err != nil {
		return commandRequest, err
	}

	return commandRequest, nil
}
