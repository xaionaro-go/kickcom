package kickcom

import "time"

// UserV1 is a representation of an user in API v1.
type UserV1 struct {
	ID              uint64    `json:"id"`
	Username        string    `json:"username"`
	AgreedToTerms   bool      `json:"agreed_to_terms"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
	Bio             string    `json:"bio"`
	Country         string    `json:"country"`
	State           string    `json:"state"`
	City            string    `json:"city"`
	Instagram       string    `json:"instagram"`
	Twitter         string    `json:"twitter"`
	Youtube         string    `json:"youtube"`
	Discord         string    `json:"discord"`
	Tiktok          string    `json:"tiktok"`
	Facebook        string    `json:"facebook"`
	ProfilePic      string    `json:"profile_pic"`
}
