package grpc

import (
	"context"
	"net"
	"time"

	config "github.com/go-kratos/kratos/v2/api/kratos/config/grpc"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"

	"google.golang.org/grpc"
)

// ServerOption is gRPC server option.
type ServerOption func(o *serverOptions)

type serverOptions struct {
	network    string
	address    string
	timeout    time.Duration
	middleware middleware.Middleware
	grpcOpts   []grpc.ServerOption
}

// Network with server network.
func Network(network string) ServerOption {
	return func(o *serverOptions) {
		o.network = network
	}
}

// Address with server address.
func Address(addr string) ServerOption {
	return func(o *serverOptions) {
		o.address = addr
	}
}

// Timeout with server timeout.
func Timeout(timeout time.Duration) ServerOption {
	return func(o *serverOptions) {
		o.timeout = timeout
	}
}

// Middleware with server middleware.
func Middleware(m middleware.Middleware) ServerOption {
	return func(o *serverOptions) {
		o.middleware = m
	}
}

// Options with grpc options.
func Options(opts ...grpc.ServerOption) ServerOption {
	return func(o *serverOptions) {
		o.grpcOpts = opts
	}
}

// Apply apply server config.
func Apply(c *config.Server) ServerOption {
	return func(o *serverOptions) {
		o.network = c.Network
		o.address = c.Address
		if c.Timeout != nil {
			o.timeout = c.Timeout.AsDuration()
		}
	}
}

// Server is a gRPC server wrapper.
type Server struct {
	*grpc.Server
	opts serverOptions
}

// NewServer creates a gRPC server by options.
func NewServer(opts ...ServerOption) *Server {
	options := serverOptions{
		network: "tcp",
		address: ":9000",
		timeout: 500 * time.Millisecond,
	}
	for _, o := range opts {
		o(&options)
	}
	var grpcOpts = []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			UnaryServerInterceptor(options.middleware),
			UnaryTimeoutInterceptor(options.timeout),
		),
	}
	if len(options.grpcOpts) > 0 {
		grpcOpts = append(grpcOpts, options.grpcOpts...)
	}
	return &Server{
		opts:   options,
		Server: grpc.NewServer(grpcOpts...),
	}
}

// Start start the gRPC server.
func (s *Server) Start(ctx context.Context) error {
	lis, err := net.Listen(s.opts.network, s.opts.address)
	if err != nil {
		return err
	}
	return s.Serve(lis)
}

// Stop stop the gRPC server.
func (s *Server) Stop(ctx context.Context) error {
	s.GracefulStop()
	return nil
}

// UnaryTimeoutInterceptor returns a unary timeout interceptor.
func UnaryTimeoutInterceptor(timeout time.Duration) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return handler(ctx, req)
	}
}

// UnaryServerInterceptor returns a unary server interceptor.
func UnaryServerInterceptor(m middleware.Middleware) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = transport.NewContext(ctx, transport.Transport{Kind: "GRPC"})
		ctx = NewContext(ctx, ServerInfo{Server: info.Server, FullMethod: info.FullMethod})
		h := func(ctx context.Context, req interface{}) (interface{}, error) {
			return handler(ctx, req)
		}
		if m != nil {
			h = m(h)
		}
		reply, err := h(ctx, req)
		if err != nil {
			return nil, err
		}
		return reply, nil
	}
}
