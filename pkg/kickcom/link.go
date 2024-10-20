package kickcom

import "time"

type Link struct {
	ID          uint64    `json:"id"`
	ChannelID   uint64    `json:"channel_id"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Order       uint64    `json:"order"`
	Title       string    `json:"title"`
}
