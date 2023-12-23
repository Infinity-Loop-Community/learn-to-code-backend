package lambda

import (
	"learn-to-code/internal/interfaces/lambda/course/requestobject"
	"testing"
)

func TestReadBody(t *testing.T) {
	tests := []struct {
		name          string
		body          string
		command       requestobject.Command
		expectError   bool
		expectedError string
	}{
		{
			name:        "Valid Body",
			body:        `{"createdAt": "2023-01-01T00:00:00Z", "data": {"key": "value"}, "type": "test"}`,
			command:     requestobject.Command{},
			expectError: false,
		},
		{
			name:          "Invalid JSON Body",
			body:          `{"createdAt": "2023-01-01T00:00:00Z", "data": {"key": "value", "type": "test"}`, // Missing closing brace
			command:       requestobject.Command{},
			expectError:   true,
			expectedError: "unexpected end of JSON input",
		},
		{
			name:          "Invalid Data",
			body:          `{"createdAt": "2023-01-01T00:00:00Z", "data": {}, "type": "test"}`,
			command:       requestobject.Command{},
			expectError:   true,
			expectedError: "data validation error: data map cannot be empty",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ReadBody(tt.body, tt.command)
			if (err != nil) != tt.expectError {
				t.Errorf("ReadBody() error = %v, expectError %v", err, tt.expectError)
			}

			// If expecting an error, check if the error message is as expected
			if tt.expectError && err != nil && tt.expectedError != "" && err.Error() != tt.expectedError {
				t.Errorf("ReadBody() error = %v, expectedError %v", err, tt.expectedError)
			}
		})
	}
}
