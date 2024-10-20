package kickcom

import "time"

type Video struct {
	ID                uint64    `json:"id"`
	LiveStreamID      uint64    `json:"live_stream_id"`
	Slug              any       `json:"slug"`
	Thumb             any       `json:"thumb"`
	S3                any       `json:"s3"`
	TradingPlatformID any       `json:"trading_platform_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	UUID              string    `json:"uuid"`
	Views             uint64    `json:"views"`
	DeletedAt         any       `json:"deleted_at"`
}
