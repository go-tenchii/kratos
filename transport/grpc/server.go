package grpc

import (
	"context"
	"net"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"

	"google.golang.org/grpc"
)

// ServerOption is gRPC server option.
type ServerOption func(o *serverOptions)

type serverOptions struct {
	network    string
	address    string
	middleware middleware.Middleware
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

// Middleware with server middleware.
func Middleware(m middleware.Middleware) ServerOption {
	return func(o *serverOptions) {
		o.middleware = m
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
	}
	for _, o := range opts {
		o(&options)
	}
	return &Server{
		opts: options,
		Server: grpc.NewServer(grpc.UnaryInterceptor(
			UnaryInterceptor(options.middleware),
		)),
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

// UnaryInterceptor returns a unary server interceptor.
func UnaryInterceptor(m middleware.Middleware) grpc.UnaryServerInterceptor {
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
