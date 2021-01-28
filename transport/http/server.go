package http

import (
	"context"
	"net"
	"net/http"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/gorilla/mux"
)

// SupportPackageIsVersion1 These constants should not be referenced from any other code.
const SupportPackageIsVersion1 = true

// ServiceRegistrar wraps a single method that supports service registration.
type ServiceRegistrar interface {
	RegisterService(desc *ServiceDesc, impl interface{})
}

// ServiceDesc represents a HTTP service's specification.
type ServiceDesc struct {
	ServiceName string
	HandlerType interface{}
	Methods     []MethodDesc
	Metadata    interface{}
}

type serverMethodHandler func(srv interface{}, ctx context.Context, req *http.Request) (interface{}, error)

// MethodDesc represents a HTTP service's method specification.
type MethodDesc struct {
	Path    string
	Method  string
	Handler serverMethodHandler
}

// DecodeRequestFunc is decode request func.
type DecodeRequestFunc func(in interface{}, req *http.Request) error

// EncodeResponseFunc is encode response func.
type EncodeResponseFunc func(out interface{}, res http.ResponseWriter, req *http.Request) error

// EncodeErrorFunc is encode error func.
type EncodeErrorFunc func(err error, res http.ResponseWriter, req *http.Request)

// RecoveryHandlerFunc is recovery handler func.
type RecoveryHandlerFunc func(ctx context.Context, req, err interface{}) error

// ServerOption is HTTP server option.
type ServerOption func(*serverOptions)

type serverOptions struct {
	network         string
	address         string
	requestDecoder  DecodeRequestFunc
	responseEncoder EncodeResponseFunc
	errorEncoder    EncodeErrorFunc
	middleware      middleware.Middleware
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

// Middleware with server middleware option.
func Middleware(m middleware.Middleware) ServerOption {
	return func(s *serverOptions) {
		s.middleware = m
	}
}

// RequestDecoder with decode request option.
func RequestDecoder(fn EncodeErrorFunc) ServerOption {
	return func(s *serverOptions) {
		s.errorEncoder = fn
	}
}

// ResponseEncoder with response handler option.
func ResponseEncoder(fn EncodeResponseFunc) ServerOption {
	return func(s *serverOptions) {
		s.responseEncoder = fn
	}
}

// ErrorEncoder with error handler option.
func ErrorEncoder(fn EncodeErrorFunc) ServerOption {
	return func(s *serverOptions) {
		s.errorEncoder = fn
	}
}

// Server is a HTTP server wrapper.
type Server struct {
	*http.Server
	router *mux.Router
	opts   serverOptions
}

// NewServer creates a HTTP server by options.
func NewServer(opts ...ServerOption) *Server {
	options := serverOptions{
		network:         "tcp",
		address:         ":8000",
		requestDecoder:  DefaultRequestDecoder,
		responseEncoder: DefaultResponseEncoder,
		errorEncoder:    DefaultErrorEncoder,
	}
	for _, o := range opts {
		o(&options)
	}
	router := mux.NewRouter()
	return &Server{
		opts:   options,
		router: router,
		Server: &http.Server{
			Handler: router,
		},
	}
}

// Handle registers a new route with a matcher for the URL path.
func (s *Server) Handle(path string, handler http.Handler) {
	s.router.Handle(path, handler)
}

// HandleFunc registers a new route with a matcher for the URL path.
func (s *Server) HandleFunc(path string, h func(http.ResponseWriter, *http.Request)) {
	s.router.HandleFunc(path, h)
}

// ServeHTTP should write reply headers and data to the ResponseWriter and then return.
func (s *Server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx := transport.NewContext(req.Context(), transport.Transport{Kind: "HTTP"})
	ctx = NewContext(ctx, ServerInfo{Request: req, Response: res})
	s.router.ServeHTTP(res, req.WithContext(ctx))
}

// Start start the HTTP server.
func (s *Server) Start(ctx context.Context) error {
	lis, err := net.Listen(s.opts.network, s.opts.address)
	if err != nil {
		return err
	}
	return s.Serve(lis)
}

// Stop stop the HTTP server.
func (s *Server) Stop(ctx context.Context) error {
	return s.Shutdown(ctx)
}

// RegisterService registers a service and its implementation to the HTTP server.
func (s *Server) RegisterService(sd *ServiceDesc, ss interface{}) {
	for _, method := range sd.Methods {
		s.registerHandle(ss, method)
	}
}

func (s *Server) registerHandle(srv interface{}, md MethodDesc) {
	s.router.HandleFunc(md.Path, func(res http.ResponseWriter, req *http.Request) {
		handler := func(ctx context.Context, in interface{}) (interface{}, error) {
			return md.Handler(srv, ctx, req)
		}
		if s.opts.middleware != nil {
			handler = s.opts.middleware(handler)
		}
		reply, err := handler(req.Context(), req)
		if err != nil {
			s.opts.errorEncoder(err, res, req)
			return
		}

		if err := s.opts.responseEncoder(reply, res, req); err != nil {
			s.opts.errorEncoder(err, res, req)
			return
		}

	}).Methods(md.Method)
}
