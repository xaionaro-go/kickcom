package kickcom

type CategoryV2 struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	Slug           string   `json:"slug"`
	Tags           []string `json:"tags"`
	ParentCategory struct {
		ID   int    `json:"id"`
		Slug string `json:"slug"`
	} `json:"parent_category"`
}
