package http

import (
	"context"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	config "github.com/go-kratos/kratos/v2/api/kratos/config/http"
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/errors"
)

// ClientOption is HTTP client option.
type ClientOption func(*Client)

// WithTimeout with client request timeout.
func WithTimeout(d time.Duration) ClientOption {
	return func(o *Client) {
		o.timeout = d
	}
}

// WithKeepAlive with client keepavlie.
func WithKeepAlive(d time.Duration) ClientOption {
	return func(o *Client) {
		o.keepAlive = d
	}
}

// WithMaxIdleConns with client max idle conns.
func WithMaxIdleConns(n int) ClientOption {
	return func(o *Client) {
		o.maxIdleConns = n
	}
}

// WithUserAgent with client user agent.
func WithUserAgent(ua string) ClientOption {
	return func(o *Client) {
		o.userAgent = ua
	}
}

// WithApply apply client config.
func WithApply(c *config.Client) ClientOption {
	return func(o *Client) {
		if c.Timeout != nil {
			o.timeout = c.Timeout.AsDuration()
		}
		if c.Keepalive != nil {
			o.keepAlive = c.Keepalive.AsDuration()
		}
		if c.MaxIdleConns > 0 {
			o.maxIdleConns = int(c.MaxIdleConns)
		}
		if c.UserAgent != "" {
			o.userAgent = c.UserAgent
		}
	}
}

// Client is a HTTP transport client.
type Client struct {
	base         http.RoundTripper
	timeout      time.Duration
	keepAlive    time.Duration
	maxIdleConns int
	userAgent    string
}

// NewClient new a HTTP transport client.
func NewClient(opts ...ClientOption) (*http.Client, error) {
	client := &Client{
		timeout:      500 * time.Millisecond,
		keepAlive:    30 * time.Second,
		maxIdleConns: 100,
	}
	for _, o := range opts {
		o(client)
	}
	client.base = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   client.timeout,
			KeepAlive: client.keepAlive,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          client.maxIdleConns,
		MaxIdleConnsPerHost:   client.maxIdleConns,
		IdleConnTimeout:       client.keepAlive,
		TLSHandshakeTimeout:   client.timeout,
		ExpectContinueTimeout: client.timeout,
	}
	return &http.Client{Transport: client}, nil
}

// RoundTrip is transport round trip.
func (c *Client) RoundTrip(req *http.Request) (*http.Response, error) {
	if c.userAgent != "" && req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", c.userAgent)
	}
	ctx, cancel := context.WithTimeout(req.Context(), c.timeout)
	defer cancel()
	res, err := c.base.RoundTrip(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CheckResponse returns an error (of type *Error) if the response
// status code is not 2xx.
func CheckResponse(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	contentType := res.Header.Get("content-type")
	codec := encoding.GetCodec(contentSubtype(contentType))
	if codec == nil {
		return errors.Unknown("Unknown", "unknown contentType: %s", contentType)
	}
	se := &errors.StatusError{}
	if err := codec.Unmarshal(data, se); err != nil {
		return err
	}
	return se
}

// DecodeResponse decodes the body of res into target. If there is no body, target is unchanged.
func DecodeResponse(res *http.Response, v interface{}) error {
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	contentType := res.Header.Get("content-type")
	codec := encoding.GetCodec(contentSubtype(contentType))
	if codec == nil {
		return errors.Unknown("Unknown", "unknown contentType: %s", contentType)
	}
	return codec.Unmarshal(data, v)
}
