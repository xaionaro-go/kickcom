package kickcom

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func (k *Kick) GetChannelV1(
	ctx context.Context,
	channel string,
) (*ChannelV1, error) {
	return Request[ChannelV1](
		ctx,
		k,
		http.MethodGet,
		fmt.Sprintf("api/v1/channels/%s", channel),
		nil,
		NoBody,
	)
}

type ChannelV1 struct {
	ID                  uint64            `json:"id"`
	UserID              uint64            `json:"user_id"`
	Slug                string            `json:"slug"`
	IsBanned            bool              `json:"is_banned"`
	PlaybackURL         string            `json:"playback_url"`
	NameUpdatedAt       string            `json:"name_updated_at"`
	VodEnabled          bool              `json:"vod_enabled"`
	SubscriptionEnabled bool              `json:"subscription_enabled"`
	FollowersCount      uint64            `json:"followersCount"`
	SubscriberBadges    []SubscriberBadge `json:"subscriber_badges"`
	BannerImage         BannerImage       `json:"banner_image"`
	RecentCategories    []Category        `json:"recent_categories"`
	Livestream          Livestream        `json:"livestream"`
	Role                any               `json:"role"`
	Muted               bool              `json:"muted"`
	FollowerBadges      []any             `json:"follower_badges"`
	OfflineBannerImage  Image             `json:"offline_banner_image"`
	CanHost             bool              `json:"can_host"`
	User                User              `json:"user"`
	Chatroom            ChatroomV1Short   `json:"chatroom"`
	AscendingLinks      []Link            `json:"ascending_links"`
	Plan                any               `json:"plan"`
	PreviousLivestreams []Livestream      `json:"previous_livestreams"`
	Verified            struct {
		ID        uint64    `json:"id"`
		ChannelID uint64    `json:"channel_id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"verified"`
	Media []MediaAsset `json:"media"`
}
