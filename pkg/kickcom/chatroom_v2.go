package kickcom

import (
	"context"
	"fmt"
	"net/http"
)

func (k *Kick) GetChatroomV2(
	ctx context.Context,
	channelSlug string,
) (*ChatroomV2, error) {
	return Request[ChatroomV2](
		ctx,
		k,
		http.MethodGet,
		fmt.Sprintf("/api/v2/channels/%s/chatroom", channelSlug),
		nil,
		NoBody,
	)
}

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
