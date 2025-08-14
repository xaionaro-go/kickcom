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
	"github.com/xaionaro-go/xsync"
)

// Kick is a client to Kick.com API.
type Kick struct {
	Client *http.Client
	Mutex  xsync.Mutex
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
	req REQUEST,
) (_ret *REPLY, _err error) {
	logger.Debugf(ctx, "Request: %s %s: %#+v, %#+v", httpMethod, route, uriValues, req)
	defer func() {
		logger.Debugf(ctx, "Reply: %#+v %v", _ret, _err)
	}()
	return xsync.DoR2(ctx, &k.Mutex, func() (*REPLY, error) {
		return request[REPLY](ctx, k, httpMethod, route, routeVars, uriValues, req)
	})
}

func request[REPLY any, REQUEST any](
	ctx context.Context,
	k *Kick,
	httpMethod string,
	route Route,
	routeVars RouteVars,
	uriValues url.Values,
	request REQUEST,
) (_ret *REPLY, _err error) {
	dstURL := GetURL(route, routeVars)
	dstURL.RawQuery = uriValues.Encode()

	var reqBody io.Reader
	if _, ok := any(request).(noBodyT); !ok {
		b, err := json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("unable to JSON-ize the request %#+v: %w", request, err)
		}

		reqBody = io.NopCloser(bytes.NewReader(b))
	}

	req, err := http.NewRequestWithContext(ctx, httpMethod, dstURL.String(), reqBody)
	if err != nil {
		return nil, fmt.Errorf("unable to create a request for %s: %w", dstURL.String(), err)
	}
	req.Header.Set("User-Agent", "application/json")

	logger.Tracef(ctx, "resulting request: %s", spew.Sdump(req))
	resp, err := k.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to make request %#+v: %w", spew.Sdump(req), err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read the response body: %w", err)
	}
	logger.Tracef(ctx, "response body: <%s>", respBody)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("the received status code is not 200: %d; body: '%s'", resp.StatusCode, respBody)
	}

	var result REPLY
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return nil, fmt.Errorf("unable to un-JSON-ize the response: %w: '%s'", err, respBody)
	}

	return &result, nil
}
