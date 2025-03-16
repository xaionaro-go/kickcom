package kickcom

import (
	"context"
	"net/http"
)

// CategoryV1Short is a representation of a category essentials in the API v1.
type CategoryV1Short struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Icon     string `json:"icon,omitempty"`
	IsMature bool   `json:"is_mature"`
}

// CategoryV1 is a full representation of a category in the API v1.
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

// CategoriesV1Reply is the response provided by GetCategoriesV1
type CategoriesV1Reply []CategoryV1Short

// GetCategoriesV1 returns the list of available categories.
func (k *Kick) GetSubcategoriesV1(
	ctx context.Context,
) (*CategoriesV1Reply, error) {
	return Request[CategoriesV1Reply](
		ctx,
		k,
		http.MethodGet,
		RouteSubcategoriesAll,
		RouteVars{},
		nil,
		NoBody,
	)
}
