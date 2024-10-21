package kickcom

import (
	"context"
	"net/http"
	"time"
)

// GetChatroomV1 returns the information about the chat root using API v1.
// For more details see the fields of ChatroomV1.
// The function is not tested.
func (k *Kick) GetChatroomV1(
	ctx context.Context,
	channelSlug string,
) (*ChatroomV1, error) {
	return Request[ChatroomV1](
		ctx,
		k,
		http.MethodGet,
		RouteChatRoomShow,
		RouteVars{"channel": channelSlug},
		nil,
		NoBody,
	)
}

// ChatroomV1Short is a representation of the chat room essentials in API v1.
type ChatroomV1Short struct {
	ID                   uint64    `json:"id"`
	ChatableType         string    `json:"chatable_type"`
	ChannelID            uint64    `json:"channel_id"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	ChatModeOld          string    `json:"chat_mode_old"`
	ChatMode             string    `json:"chat_mode"`
	SlowMode             bool      `json:"slow_mode"`
	ChatableID           uint64    `json:"chatable_id"`
	FollowersMode        bool      `json:"followers_mode"`
	SubscribersMode      bool      `json:"subscribers_mode"`
	EmotesMode           bool      `json:"emotes_mode"`
	MessageInterval      uint64    `json:"message_interval"`
	FollowingMinDuration uint64    `json:"following_min_duration"`
}

// ChatroomV1 is the full representation of a chat room in API v1.
type ChatroomV1 struct {
	ID                  int             `json:"id"`
	UserID              int             `json:"user_id"`
	Slug                string          `json:"slug"`
	IsBanned            bool            `json:"is_banned"`
	PlaybackURL         string          `json:"playback_url"`
	NameUpdatedAt       any             `json:"name_updated_at"`
	VodEnabled          bool            `json:"vod_enabled"`
	SubscriptionEnabled bool            `json:"subscription_enabled"`
	Role                any             `json:"role"`
	FollowerBadges      []any           `json:"follower_badges"`
	MutedUsers          []any           `json:"muted_users"`
	CanHost             bool            `json:"can_host"`
	Chatroom            ChatroomV1Short `json:"chatroom"`
	Emotes              []EmoteV1       `json:"emotes"`
}
