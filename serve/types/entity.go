package types

import (
	"encoding/json"
	"notify-api/push/types"
	"time"
)

type Message struct {
	ID        string         `json:"id"`
	UserID    string         `json:"user_id"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	Long      string         `json:"long"`
	Priority  types.Priority `json:"priority"`
	CreatedAt time.Time      `json:"created_at"`
}

func (m Message) MarshalJSON() ([]byte, error) {
	type Alias Message
	return json.Marshal(&struct {
		*Alias
		CreatedAt string `json:"created_at"`
	}{
		Alias:     (*Alias)(&m),
		CreatedAt: m.CreatedAt.Format(time.RFC3339),
	})
}
