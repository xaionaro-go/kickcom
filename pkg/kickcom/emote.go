package kickcom

type Emote struct {
	ID              int    `json:"id"`
	ChannelID       int    `json:"channel_id"`
	Name            string `json:"name"`
	SubscribersOnly bool   `json:"subscribers_only"`
	Image           any    `json:"image"`
}
