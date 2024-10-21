package kickcom

// CategoryV1Short is a representation of a category essentials in the API v1.
type CategoryV1Short struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Icon string `json:"icon"`
}

// CategoryV1Short is a full representation of a category in the API v1.
type CategoryV1 struct {
	ID          uint64          `json:"id"`
	CategoryID  uint64          `json:"category_id"`
	Name        string          `json:"name"`
	Slug        string          `json:"slug"`
	Tags        []string        `json:"tags"`
	Description any             `json:"description"`
	DeletedAt   any             `json:"deleted_at"`
	Viewers     uint64          `json:"viewers"`
	Category    CategoryV1Short `json:"category"`
}
