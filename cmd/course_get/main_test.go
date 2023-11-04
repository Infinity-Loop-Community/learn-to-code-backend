package main

import (
	"encoding/json"
	err2 "learn-to-code/internal/infrastructure/go/util/err"
	"learn-to-code/internal/infrastructure/inmemory"
	"learn-to-code/internal/infrastructure/local"
	course "learn-to-code/internal/interfaces/lambda/course"
	"testing"

	"github.com/aws/aws-lambda-go/events"

	"github.com/yalp/jsonpath"
)

func TestGetCourseLambda_Returns200(t *testing.T) {

	environmentCreator := local.NewEnvironmentCreator()
	handlerResponse := environmentCreator.ExecuteLambdaHandler(course.NewLambdaHandler(environmentCreator.Cfg).HandleRequest)

	if handlerResponse.StatusCode != 200 {
		t.Fatalf("lambda did not succeed, status code: %v", handlerResponse.StatusCode)
	}
}

func TestGetCourseLambda_ContainsCourseData(t *testing.T) {

	environmentCreator := local.NewEnvironmentCreator()
	handlerResponse := environmentCreator.ExecuteLambdaHandler(course.NewLambdaHandler(environmentCreator.Cfg).HandleRequest)

	path := "$.id"
	courseID := getJSONPathValue(handlerResponse, path)
	if courseID != inmemory.CourseID {
		t.Fatalf("expected course '%s' not found: '%s'", inmemory.CourseID, courseID)
	}
}

func getJSONPathValue(handlerResponse events.APIGatewayProxyResponse, path string) interface{} {
	raw := []byte(handlerResponse.Body)

	var data interface{}
	json.Unmarshal(raw, &data)

	courseID := err2.PanicIfError1(jsonpath.Read(data, path))
	return courseID
}
