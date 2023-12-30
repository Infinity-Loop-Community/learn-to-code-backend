package main

import (
	"learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/inmemory"
	"learn-to-code/internal/infrastructure/local"
	"learn-to-code/internal/infrastructure/testing/json"
	course "learn-to-code/internal/interfaces/lambda/course"
	"testing"
)

func TestGetCourseLambda_Returns200(t *testing.T) {

	environmentCreator := local.NewEnvironmentCreator(config.Test)
	handlerResponse := environmentCreator.ExecuteLambdaHandler(course.NewLambdaHandler)

	if handlerResponse.StatusCode != 200 {
		t.Fatalf("lambda did not succeed, status code: %v", handlerResponse.StatusCode)
	}
}

func TestGetCourseLambda_ContainsCourseData(t *testing.T) {

	environmentCreator := local.NewEnvironmentCreator(config.Test)
	defer environmentCreator.Terminate()

	handlerResponse := environmentCreator.ExecuteLambdaHandler(course.NewLambdaHandler)

	path := "$.id"
	courseID := json.GetJSONPathValue(handlerResponse, path)
	if courseID != inmemory.CourseIDFrontendDevelopment {
		t.Fatalf("expected course '%s' not found: '%s'", inmemory.CourseIDFrontendDevelopment, courseID)
	}
}
