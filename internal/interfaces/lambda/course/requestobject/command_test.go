package requestobject

import (
	"testing"
	"time"
)

func TestCommand_Validate(t *testing.T) {
	tests := []struct {
		name      string
		command   Command
		wantError bool
	}{
		{
			name: "Valid Command",
			command: Command{
				CreatedAt: time.Now(),
				Data:      map[string]interface{}{"key": "value"},
				Type:      "example",
			},
			wantError: false,
		},
		{
			name: "Zero CreatedAt",
			command: Command{
				CreatedAt: time.Time{},
				Data:      map[string]interface{}{"key": "value"},
				Type:      "example",
			},
			wantError: true,
		},
		{
			name: "Empty Type",
			command: Command{
				CreatedAt: time.Now(),
				Data:      map[string]interface{}{"key": "value"},
				Type:      "",
			},
			wantError: true,
		},
		{
			name: "Nil Data",
			command: Command{
				CreatedAt: time.Now(),
				Type:      "example",
			},
			wantError: true,
		},
		{
			name: "Empty Data Field",
			command: Command{
				CreatedAt: time.Now(),
				Data:      map[string]interface{}{"key": ""},
				Type:      "example",
			},
			wantError: true,
		},
		{
			name: "Empty Data Map",
			command: Command{
				CreatedAt: time.Now(),
				Data:      map[string]interface{}{},
				Type:      "example",
			},
			wantError: true,
		},
		{
			name: "Data with Nested Empty Field",
			command: Command{
				CreatedAt: time.Now(),
				Data:      map[string]interface{}{"nested": map[string]interface{}{"key": ""}},
				Type:      "example",
			},
			wantError: true,
		},
		{
			name: "Data with Zero Integer",
			command: Command{
				CreatedAt: time.Now(),
				Data:      map[string]interface{}{"number": 0},
				Type:      "example",
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.command.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("Command.Validate() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}
