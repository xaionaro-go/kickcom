package kickcom

import (
	"context"
	"net/http"
	"time"
)

// GetChannelV1 returns channel information given channel slug using API v1.
func (k *Kick) GetChannelV1(
	ctx context.Context,
	channel string,
) (*ChannelV1, error) {
	return Request[ChannelV1](
		ctx,
		k,
		http.MethodGet,
		RouteChannelsShow,
		RouteVars{"channel": channel},
		nil,
		NoBody,
	)
}

// ChannelV1 is a representation of a channel in the API v1.
type ChannelV1 struct {
	ID                  uint64              `json:"id"`
	UserID              uint64              `json:"user_id"`
	Slug                string              `json:"slug"`
	IsBanned            bool                `json:"is_banned"`
	PlaybackURL         string              `json:"playback_url"`
	NameUpdatedAt       string              `json:"name_updated_at"`
	VodEnabled          bool                `json:"vod_enabled"`
	SubscriptionEnabled bool                `json:"subscription_enabled"`
	FollowersCount      uint64              `json:"followersCount"`
	SubscriberBadges    []SubscriberBadgeV1 `json:"subscriber_badges"`
	BannerImage         BannerImageV1       `json:"banner_image"`
	RecentCategories    []CategoryV1        `json:"recent_categories"`
	Livestream          LivestreamV1        `json:"livestream"`
	Role                any                 `json:"role"`
	Muted               bool                `json:"muted"`
	FollowerBadges      []any               `json:"follower_badges"`
	OfflineBannerImage  ImageV1             `json:"offline_banner_image"`
	CanHost             bool                `json:"can_host"`
	User                UserV1              `json:"user"`
	Chatroom            ChatroomV1Short     `json:"chatroom"`
	AscendingLinks      []LinkV1            `json:"ascending_links"`
	Plan                any                 `json:"plan"`
	PreviousLivestreams []LivestreamV1      `json:"previous_livestreams"`
	Verified            struct {
		ID        uint64    `json:"id"`
		ChannelID uint64    `json:"channel_id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"verified"`
	Media []MediaAssetV1 `json:"media"`
}
