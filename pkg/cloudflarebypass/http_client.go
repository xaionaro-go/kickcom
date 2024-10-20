package cloudflarebypass

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/Danny-Dasilva/CycleTLS/cycletls"
	"github.com/RomainMichau/cloudscraper_go/cloudscraper"
)

type RoundTripper struct {
	*cloudscraper.CloudScrapper
}

var _ http.RoundTripper = (*RoundTripper)(nil)

func New() (*RoundTripper, error) {
	c, err := cloudscraper.Init(false, false)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize a 'cloudscraper': %w", err)
	}

	return &RoundTripper{
		CloudScrapper: c,
	}, nil
}

func (r *RoundTripper) RoundTrip(
	req *http.Request,
) (*http.Response, error) {
	var (
		body []byte
		err  error
	)
	if req.Body != nil {
		body, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, fmt.Errorf("unable to read the request body: %w", err)
		}
	}

	headerMap := map[string]string{}
	for k, v := range req.Header {
		if len(v) != 1 {
			return nil, fmt.Errorf("the case with multiple HTTP headers is not supported: %d != 1", len(v))
		}
		headerMap[k] = v[0]
	}

	var resp cycletls.Response
	switch req.Method {
	case http.MethodGet:
		resp, err = r.CloudScrapper.Get(req.URL.String(), headerMap, string(body))
	case http.MethodPost:
		resp, err = r.CloudScrapper.Post(req.URL.String(), headerMap, string(body))
	default:
		resp, err = r.CloudScrapper.Do(req.URL.String(), cycletls.Options{}, req.Method)
	}
	if err != nil {
		return nil, fmt.Errorf("unable to query: %w", err)
	}

	respHeader := http.Header{}
	for k, v := range resp.Headers {
		respHeader[k] = []string{v}
	}

	return &http.Response{
		Status:        fmt.Sprintf("%d %s", resp.Status, http.StatusText(resp.Status)),
		StatusCode:    resp.Status,
		Proto:         "HTTP/1.0",
		ProtoMajor:    1,
		ProtoMinor:    0,
		Header:        respHeader,
		Body:          io.NopCloser(bytes.NewReader([]byte(resp.Body))),
		ContentLength: int64(len(resp.Body)),
		Request:       req,
	}, nil
}
