package event

import (
	"time"
)

type JoinedQuiz struct {
	Id   string
	Time time.Time
}
