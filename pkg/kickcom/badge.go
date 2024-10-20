package kickcom

type Badge struct {
	Type   string `json:"type"`
	Text   string `json:"text"`
	Active bool   `json:"active"`
	Count  int    `json:"count,omitempty"`
}
