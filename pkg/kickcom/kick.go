package kickcom

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/facebookincubator/go-belt/tool/logger"
	"github.com/xaionaro-go/kickcom/pkg/cloudflarebypass"
)

type Kick struct {
	*http.Client
	BaseURL *url.URL
}

func New() (*Kick, error) {
	const baseURLString = "https://kick.com"
	baseURL, err := url.Parse(baseURLString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse URL '%s': %w", baseURLString, err)
	}

	cj, err := cookiejar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize a cookie jar: %w", err)
	}

	rt, err := cloudflarebypass.New()
	if err != nil {
		return nil, fmt.Errorf("unable to initialize a cloudflare-bypass: %w", err)
	}

	return &Kick{
		Client: &http.Client{
			Transport: rt,
			Jar:       cj,
		},
		BaseURL: baseURL,
	}, nil
}

type noBodyT struct{}

var NoBody noBodyT

func Request[REPLY any, REQUEST any](
	ctx context.Context,
	k *Kick,
	httpMethod string,
	path string,
	uriValues url.Values,
	request REQUEST,
) (_ret *REPLY, _err error) {
	logger.Debugf(ctx, "Request: %s %s: %#+v", httpMethod, path, request)
	defer func() {
		logger.Debugf(ctx, "Reply: %#+v %v", _ret, _err)
	}()

	dstURL := ptr(*k.BaseURL)
	if !strings.HasSuffix(dstURL.Path, "/") {
		dstURL.Path += "/"
	}
	dstURL.Path += path
	dstURL.RawQuery = uriValues.Encode()

	req := &http.Request{
		URL:    dstURL,
		Method: httpMethod,
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
	}

	var b []byte
	if _, ok := any(request).(noBodyT); !ok {
		var err error
		b, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("unable to JSON-ize the request %#+v: %w", request, err)
		}

		req.Body = io.NopCloser(bytes.NewReader(b))
	}

	logger.Tracef(ctx, "resulting request: %s", spew.Sdump(req))
	resp, err := k.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to make request %#+v: %w", spew.Sdump(req), err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read the response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("the received status code is not 200: %d; body: '%s'", resp.StatusCode, body)
	}

	var result REPLY
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("unable to un-JSON-ize the response: %w: '%s'", err, body)
	}

	return &result, nil
}
