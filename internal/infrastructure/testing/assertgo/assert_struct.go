package assertgo

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type assertStruct struct {
	t                   *testing.T
	source              interface{}
	typeKeyManipulation func(string) string
}

func NewAssertion(t *testing.T, source interface{}) Assertion {
	return &assertStruct{t: t, source: source}
}

func (a *assertStruct) WithTypeKeyManipulation(manipulationFunc func(string) string) Assertion {
	a.typeKeyManipulation = manipulationFunc
	return a
}

func (a *assertStruct) IsEqualTo(target interface{}) {
	sourceMap := a.convertToMap(a.source)
	targetMap := a.convertToMap(target)

	if a.typeKeyManipulation != nil {
		sourceMap = deepTransformKeys(sourceMap, a.typeKeyManipulation)
		targetMap = deepTransformKeys(targetMap, a.typeKeyManipulation)
	}

	a.compareMaps(sourceMap, targetMap)
}

func (a *assertStruct) convertToMap(value interface{}) map[string]interface{} {
	bytes, err := json.Marshal(value)
	if err != nil {
		a.t.Fatalf("Error marshaling: %v", err)
	}
	var resultMap map[string]interface{}
	if err = json.Unmarshal(bytes, &resultMap); err != nil {
		a.t.Fatalf("Error unmarshaling: %v", err)
	}
	return resultMap
}

func (a *assertStruct) compareMaps(source, target map[string]interface{}) {
	diff := cmp.Diff(source, target)
	if diff != "" {
		a.t.Fatalf("Mismatch (-want +got):\n%s", diff)
	}
}

func deepTransformKeys(original map[string]interface{}, transformFunc func(string) string) map[string]interface{} {
	transformed := make(map[string]interface{})
	for key, value := range original {
		transformedKey := transformFunc(key)
		switch v := value.(type) {
		case map[string]interface{}:
			transformed[transformedKey] = handleMap(v, transformFunc)
		case []interface{}:
			transformed[transformedKey] = handleArray(v, transformFunc)
		default:
			transformed[transformedKey] = value
		}
	}
	return transformed
}

func handleMap(original map[string]interface{}, transformFunc func(string) string) map[string]interface{} {
	return deepTransformKeys(original, transformFunc)
}

func handleArray(original []interface{}, transformFunc func(string) string) []interface{} {
	var newArray []interface{}
	for _, item := range original {
		switch itemValue := item.(type) {
		case map[string]interface{}:
			newArray = append(newArray, handleMap(itemValue, transformFunc))
		default:
			newArray = append(newArray, item)
		}
	}
	return newArray
}
