package kickcom

// SubscriberBadgeV1 is a representation of a subscriber badge in API v1.
type SubscriberBadgeV1 struct {
	ID         uint64  `json:"id"`
	ChannelID  uint64  `json:"channel_id"`
	Months     uint64  `json:"months"`
	BadgeImage ImageV1 `json:"badge_image"`
}
