package requestobject

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Command struct {
	CreatedAt time.Time `json:"createdAt"`
	Data      any       `json:"data"`
	Type      string    `json:"type"`
}

func (c Command) Validate() error {
	if c.CreatedAt.IsZero() {
		return errors.New("created at cannot be the zero value")
	}

	if c.Type == "" {
		return errors.New("type cannot be empty")
	}

	if c.Data == nil {
		return errors.New("data cannot be nil")
	}

	// Validate Data field
	if err := validateData(c.Data); err != nil {
		return fmt.Errorf("data validation error: %w", err)
	}

	return nil
}

func validateData(data any) error {
	// Marshal and then unmarshal the data into a map for inspection
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return errors.New("unable to marshal data")
	}

	var dataMap map[string]interface{}
	if err := json.Unmarshal(dataBytes, &dataMap); err != nil {
		return errors.New("unable to unmarshal data into a map")
	}

	err = checkIfEmptyMap(dataMap)
	if err != nil {
		return err
	}

	err = checkIfEmptyEachField(dataMap)
	if err != nil {
		return err
	}

	return nil
}

func checkIfEmptyEachField(dataMap map[string]interface{}) error {
	for key, value := range dataMap {
		if isEmpty(value) {
			return fmt.Errorf("field '%s' is empty", key)
		}
	}
	return nil
}

func checkIfEmptyMap(dataMap map[string]interface{}) error {
	if len(dataMap) == 0 {
		return errors.New("data map cannot be empty")
	}
	return nil
}

func isEmpty(value interface{}) bool {
	if value == nil {
		return true
	}

	switch v := value.(type) {
	case string:
		return v == ""
	case int, int32, int64, float32, float64:
		return v == 0
	case []interface{}:
		return len(v) == 0
	case map[string]interface{}:
		for _, val := range v {
			if isEmpty(val) {
				return true
			}
		}
	}

	return false
}
