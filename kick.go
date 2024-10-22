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

	"github.com/davecgh/go-spew/spew"
	"github.com/facebookincubator/go-belt/tool/logger"
	"github.com/xaionaro-go/kickcom/cloudflarebypass"
)

// Kick is a client to Kick.com API.
type Kick struct {
	*http.Client
}

// New returns a new instance of Kick.
func New() (*Kick, error) {
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
	}, nil
}

type noBodyT struct{}

// NoBody is a request that could be used in function `Request` to signify
// that there is no request body.
var NoBody noBodyT

// Request performs a request to Kick's backend.
func Request[REPLY any, REQUEST any](
	ctx context.Context,
	k *Kick,
	httpMethod string,
	route Route,
	routeVars RouteVars,
	uriValues url.Values,
	request REQUEST,
) (_ret *REPLY, _err error) {
	logger.Debugf(ctx, "Request: %s %s: %#+v, %#+v", httpMethod, route, uriValues, request)
	defer func() {
		logger.Debugf(ctx, "Reply: %#+v %v", _ret, _err)
	}()

	dstURL := GetURL(route, routeVars)
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
	logger.Tracef(ctx, "response body: <%s>", body)

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
