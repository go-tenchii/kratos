package http

import (
	"context"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/errors"
)

// ClientOption is HTTP client option.
type ClientOption func(*Client)

// ClientDecodeErrorFunc is client error decoder.
type ClientDecodeErrorFunc func(res *http.Response) error

// ClientTimeout with client request timeout.
func ClientTimeout(d time.Duration) ClientOption {
	return func(c *Client) {
		c.timeout = d
	}
}

// ClientKeepAlive with client keepavlie.
func ClientKeepAlive(d time.Duration) ClientOption {
	return func(c *Client) {
		c.keepAlive = d
	}
}

// ClientMaxIdleConns with client max idle conns.
func ClientMaxIdleConns(n int) ClientOption {
	return func(c *Client) {
		c.maxIdleConns = n
	}
}

// ClientUserAgent with client user agent.
func ClientUserAgent(ua string) ClientOption {
	return func(c *Client) {
		c.userAgent = ua
	}
}

// ClientErrorDecoder with client error decoder.
func ClientErrorDecoder(d ClientDecodeErrorFunc) ClientOption {
	return func(c *Client) {
		c.errorDecoder = d
	}
}

// Client is a HTTP transport client.
type Client struct {
	base         http.RoundTripper
	timeout      time.Duration
	keepAlive    time.Duration
	maxIdleConns int
	userAgent    string
	errorDecoder ClientDecodeErrorFunc
}

// NewClient new a HTTP transport client.
func NewClient(opts ...ClientOption) (*http.Client, error) {
	client := &Client{
		timeout:      500 * time.Millisecond,
		keepAlive:    30 * time.Second,
		maxIdleConns: 100,
		errorDecoder: CheckResponse,
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
func (c *Client) RoundTrip(req *http.Request) (res *http.Response, err error) {
	if c.userAgent != "" && req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", c.userAgent)
	}

	ctx, cancel := context.WithTimeout(req.Context(), c.timeout)
	defer cancel()
	if res, err = c.base.RoundTrip(req.WithContext(ctx)); err != nil {
		return nil, err
	}

	if err = c.errorDecoder(res); err != nil {
		return nil, err
	}
	return
}

// CheckResponse returns an error (of type *Error) if the response
// status code is not 2xx.
func CheckResponse(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}
	defer res.Body.Close()
	slurp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	contentType := res.Header.Get("content-type")
	codec := encoding.GetCodec(contentSubtype(contentType))
	if codec == nil {
		return errors.Unknown("Unknown", "unknown contentType: %s", contentType)
	}
	code, ok := statusMapping[res.StatusCode]
	if !ok {
		code = 2
	}
	se := &errors.StatusError{Code: code}
	if err := codec.Unmarshal(slurp, se); err != nil {
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
