package kickcom

// BadgeV2 is a representation of a badge in the API v2.
type BadgeV2 struct {
	Type   string `json:"type"`
	Text   string `json:"text"`
	Active bool   `json:"active"`
	Count  int    `json:"count,omitempty"`
}
