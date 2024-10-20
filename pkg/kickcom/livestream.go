package kickcom

type Livestream struct {
	ID            uint64 `json:"id"`
	Slug          string `json:"slug"`
	ChannelID     uint64 `json:"channel_id"`
	CreatedAt     string `json:"created_at"`
	SessionTitle  string `json:"session_title"`
	IsLive        bool   `json:"is_live"`
	RiskLevelID   any    `json:"risk_level_id"`
	StartTime     string `json:"start_time"`
	Source        any    `json:"source"`
	TwitchChannel any    `json:"twitch_channel"`
	Duration      uint64 `json:"duration"`
	Language      string `json:"language"`
	IsMature      bool   `json:"is_mature"`
	ViewerCount   uint64 `json:"viewer_count"`
	Thumbnail     struct {
		Src    string `json:"src"`
		Srcset string `json:"srcset"`
	} `json:"thumbnail"`
	Views      uint64     `json:"views"`
	Tags       []any      `json:"tags"`
	Categories []Category `json:"categories"`
	Video      Video      `json:"video"`
}
