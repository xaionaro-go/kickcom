package kickcom

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

//go:embed resources/endpoints.json
var endpointsBytes []byte

type routeInfo struct {
	URI      string            `json:"uri"`
	Methods  []string          `json:"methods"`
	Bindings map[string]string `json:"bindings"`
}

type endpointsT struct {
	BaseURL  string   `json:"url"`
	Port     *uint16  `json:"port"`
	Defaults struct{} `json:"defaults"`
	Routes   map[Route]routeInfo
}

var endpoints endpointsT
var baseURL *url.URL

func init() {
	err := json.Unmarshal(endpointsBytes, &endpoints)
	if err != nil {
		panic(err)
	}

	baseURL, err = url.Parse(endpoints.BaseURL)
	if err != nil {
		panic(fmt.Errorf("unable to parse URL '%s': %w", endpoints.BaseURL, err))
	}

	if !strings.HasSuffix(baseURL.Path, "/") {
		baseURL.Path += "/"
	}
}

// RouteVars are the substitutions in an URI defined in ./resources/endpoints.json.
type RouteVars map[string]any

// GetURL returns an URL of the requested endpoint.
func GetURL(
	route Route,
	routeVars RouteVars,
) *url.URL {
	routeInfo, ok := endpoints.Routes[route]
	if !ok {
		return nil
	}

	path := routeInfo.URI
	for k, v := range routeVars {
		path = strings.ReplaceAll(path, fmt.Sprintf("{%s}", k), fmt.Sprintf("%v", v))
	}

	dstURL := ptr(*baseURL)
	dstURL.Path += path
	return dstURL
}
