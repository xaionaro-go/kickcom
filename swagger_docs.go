package kickcom

import (
	"context"
	"net/http"
)

// TBDSwaggerGetDocs is supposed to return swagger dogs, but it does not work.
// DO NOT USE THIS FUNCTION
func (k *Kick) TBDSwaggerGetDocs(
	ctx context.Context,
) (*map[string]any, error) {
	return Request[map[string]any](
		ctx,
		k,
		http.MethodGet,
		RouteSwaggerGetDocs,
		nil,
		nil,
		NoBody,
	)
}
