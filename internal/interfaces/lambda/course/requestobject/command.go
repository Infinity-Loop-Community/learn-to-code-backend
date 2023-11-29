package requestobject

import "time"

type Command struct {
	CreatedAt time.Time `json:"createdAt"`
	Data      any       `json:"data"`
	Type      string    `json:"type"`
}
