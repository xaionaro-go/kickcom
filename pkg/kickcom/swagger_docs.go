package kickcom

import (
	"context"
	"net/http"
)

func (k *Kick) SwaggerGetDocs(
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
