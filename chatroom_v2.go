package kickcom

import (
	"context"
	"net/http"
)

// GetChatroomV2 returns the information about the chat room using API v2.
// For more details see the fields of ChatroomV2.
// The function is not tested.
func (k *Kick) GetChatroomV2(
	ctx context.Context,
	channelSlug string,
) (*ChatroomV2, error) {
	return Request[ChatroomV2](
		ctx,
		k,
		http.MethodGet,
		RouteChannelChatRoom,
		RouteVars{"channel": channelSlug},
		nil,
		NoBody,
	)
}

// ChatroomV2 contains the information about the chat room as
// it is provided by API v2.
type ChatroomV2 struct {
	ID       int `json:"id"`
	SlowMode struct {
		Enabled         bool `json:"enabled"`
		MessageInterval int  `json:"message_interval"`
	} `json:"slow_mode"`
	SubscribersMode struct {
		Enabled bool `json:"enabled"`
	} `json:"subscribers_mode"`
	FollowersMode struct {
		Enabled     bool `json:"enabled"`
		MinDuration int  `json:"min_duration"`
	} `json:"followers_mode"`
	EmotesMode struct {
		Enabled bool `json:"enabled"`
	} `json:"emotes_mode"`
	AdvancedBotProtection struct {
		Enabled       bool `json:"enabled"`
		RemainingTime int  `json:"remaining_time"`
	} `json:"advanced_bot_protection"`
	PinnedMessage   interface{} `json:"pinned_message"`
	ShowQuickEmotes struct {
		Enabled bool `json:"enabled"`
	} `json:"show_quick_emotes"`
	ShowBanners struct {
		Enabled bool `json:"enabled"`
	} `json:"show_banners"`
	GiftsEnabled struct {
		Enabled bool `json:"enabled"`
	} `json:"gifts_enabled"`
	GiftsWeekEnabled struct {
		Enabled bool `json:"enabled"`
	} `json:"gifts_week_enabled"`
	GiftsMonthEnabled struct {
		Enabled bool `json:"enabled"`
	} `json:"gifts_month_enabled"`
}
