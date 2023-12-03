package json

import (
	"encoding/json"
	"learn-to-code/internal/infrastructure/go/util/err"

	"github.com/aws/aws-lambda-go/events"
	"github.com/yalp/jsonpath"
)

func GetJSONPathValue(handlerResponse events.APIGatewayProxyResponse, path string) interface{} {
	raw := []byte(handlerResponse.Body)

	var data interface{}
	json.Unmarshal(raw, &data)

	courseID := err.PanicIfError1(jsonpath.Read(data, path))
	return courseID
}
