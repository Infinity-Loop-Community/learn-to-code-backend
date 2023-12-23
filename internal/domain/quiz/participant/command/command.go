package command

import "time"

func NewCommand(commandType string, data any, createdAt time.Time) Command {
	return Command{
		CreatedAt: createdAt,
		Data:      data,
		Type:      commandType,
	}
}

type Command struct {
	CreatedAt time.Time
	Data      any
	Type      string
}
