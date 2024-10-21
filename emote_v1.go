package kickcom

// EmoteV1 is a representation of an emote in API v1.
type EmoteV1 struct {
	ID              int    `json:"id"`
	ChannelID       int    `json:"channel_id"`
	Name            string `json:"name"`
	SubscribersOnly bool   `json:"subscribers_only"`
	Image           any    `json:"image"`
}
