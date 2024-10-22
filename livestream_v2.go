package kickcom

import (
	"context"
	"net/http"
	"time"
)

// GetLivestreamV2 returns the information about the current live stream using API v2.
func (k *Kick) GetLivestreamV2(
	ctx context.Context,
	channelSlug string,
) (*LivestreamV2Reply, error) {
	return Request[LivestreamV2Reply](
		ctx,
		k,
		http.MethodGet,
		RouteChannelLivestream,
		RouteVars{"channel": channelSlug},
		nil,
		NoBody,
	)
}

// LivestreamV2Reply is the response provided by GetLivestreamV2
type LivestreamV2Reply struct {
	Data LivestreamV2 `json:"data"`
}

// LivestreamV2 represents a live stream in API v2.
type LivestreamV2 struct {
	ID           int        `json:"id"`
	Slug         string     `json:"slug"`
	SessionTitle string     `json:"session_title"`
	CreatedAt    time.Time  `json:"created_at"`
	Language     string     `json:"language"`
	IsMature     bool       `json:"is_mature"`
	Viewers      int        `json:"viewers"`
	Category     CategoryV2 `json:"category"`
	PlaybackURL  string     `json:"playback_url"`
	Thumbnail    struct {
		Src    string `json:"src"`
		Srcset string `json:"srcset"`
	} `json:"thumbnail"`
}
