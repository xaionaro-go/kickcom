package kickcom

type SubscriberBadge struct {
	ID         uint64 `json:"id"`
	ChannelID  uint64 `json:"channel_id"`
	Months     uint64 `json:"months"`
	BadgeImage Image  `json:"badge_image"`
}
