package messages

import "time"

type BaseMessage struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Username string `json:"username"`

	CreatedAt time.Time `json:"createdAt"`
}
