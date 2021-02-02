package grpc

import (
	"context"
	"time"

	config "github.com/go-kratos/kratos/v2/api/kratos/config/grpc"
	"github.com/go-kratos/kratos/v2/middleware"

	"google.golang.org/grpc"
)

// ClientOption is gRPC client option.
type ClientOption func(o *clientOptions)

// WithContext with client context.
func WithContext(ctx context.Context) ClientOption {
	return func(c *clientOptions) {
		c.ctx = ctx
	}
}

// WithTimeout with client timeout.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *clientOptions) {
		c.timeout = timeout
	}
}

// WithInsecure with client insecure.
func WithInsecure() ClientOption {
	return func(c *clientOptions) {
		c.insecure = true
	}
}

// WithUnaryInterceptor with unary client interceptor.
func WithUnaryInterceptor(in grpc.UnaryClientInterceptor) ClientOption {
	return func(c *clientOptions) {
		c.interceptor = in
	}
}

// WithMiddleware with server middleware.
func WithMiddleware(m middleware.Middleware) ClientOption {
	return func(o *clientOptions) {
		o.middleware = m
	}
}

// WithOptions with gRPC options.
func WithOptions(opts ...grpc.DialOption) ClientOption {
	return func(o *clientOptions) {
		o.grpcOpts = opts
	}
}

// WithApply apply client config.
func WithApply(c *config.Client) ClientOption {
	return func(o *clientOptions) {
		if c.Timeout != nil {
			o.timeout = c.Timeout.AsDuration()
		}
		o.insecure = c.Insecure
	}
}

type clientOptions struct {
	ctx         context.Context
	insecure    bool
	timeout     time.Duration
	interceptor grpc.UnaryClientInterceptor
	middleware  middleware.Middleware
	grpcOpts    []grpc.DialOption
}

// NewClient new a grpc transport client.
func NewClient(target string, opts ...ClientOption) (*grpc.ClientConn, error) {
	options := clientOptions{
		ctx:      context.Background(),
		timeout:  500 * time.Millisecond,
		insecure: false,
	}
	for _, o := range opts {
		o(&options)
	}
	var grpcOpts = []grpc.DialOption{
		grpc.WithTimeout(options.timeout),
		grpc.WithChainUnaryInterceptor(
			options.interceptor,
			UnaryClientInterceptor(options.middleware),
		),
	}
	if options.insecure {
		grpcOpts = append(grpcOpts, grpc.WithInsecure())
	}
	if len(options.grpcOpts) > 0 {
		grpcOpts = append(grpcOpts, options.grpcOpts...)
	}
	return grpc.DialContext(options.ctx, target, grpcOpts...)
}

// UnaryClientInterceptor retruns a unary client interceptor.
func UnaryClientInterceptor(m middleware.Middleware) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		h := func(ctx context.Context, req interface{}) (interface{}, error) {
			return reply, invoker(ctx, method, req, reply, cc, opts...)
		}
		if m != nil {
			h = m(h)
		}
		_, err := h(ctx, req)
		return err
	}
}
