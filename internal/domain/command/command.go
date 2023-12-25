package command

import "time"

func NewCommand(commandType string, data any, createdAt time.Time) Command {
	return Command{
		CreatedAt: createdAt,
		Data:      data,
		Type:      commandType,
	}
}

// Command struct defines the structure of a command in the event-sourced system.
// It contains essential information required to process and understand the command.
type Command struct {

	// CreatedAt is the timestamp indicating when the command was created.
	CreatedAt time.Time

	// Data holds the actual payload of the command.
	Data any

	// Type is a string that signifies the type or nature of the command.
	Type string
}
