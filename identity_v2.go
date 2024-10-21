package kickcom

// IdentityV2 is a representation of a user's identity (in the chat?)
// in API v2.
type IdentityV2 struct {
	Color  string    `json:"color"`
	Badges []BadgeV2 `json:"badges"`
}
