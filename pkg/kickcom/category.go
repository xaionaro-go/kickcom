package kickcom

type CategoryShort struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Icon string `json:"icon"`
}

type Category struct {
	ID          uint64        `json:"id"`
	CategoryID  uint64        `json:"category_id"`
	Name        string        `json:"name"`
	Slug        string        `json:"slug"`
	Tags        []string      `json:"tags"`
	Description any           `json:"description"`
	DeletedAt   any           `json:"deleted_at"`
	Viewers     uint64        `json:"viewers"`
	Category    CategoryShort `json:"category"`
}
