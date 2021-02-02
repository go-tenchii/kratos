package logging

import (
	"context"
	"path"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// GRPCServer is a gRPC logging middleware.
func GRPCServer(logger log.Logger) middleware.Middleware {
	log := log.NewHelper("grpc", logger)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var (
				service string
				method  string
			)
			info, ok := grpc.FromContext(ctx)
			if ok {
				service = path.Dir(info.FullMethod)[1:]
				method = path.Base(info.FullMethod)
			}
			reply, err := handler(ctx, req)
			if err != nil {
				log.Errorw(
					"kind", "server",
					"grpc.service", service,
					"grpc.method", method,
					"grpc.code", errors.Code(err),
					"grpc.error", err.Error(),
				)
				return nil, err
			}
			log.Infow(
				"kind", "server",
				"grpc.service", service,
				"grpc.method", method,
				"grpc.code", 0,
			)
			return reply, nil
		}
	}
}

// HTTPServer is a gRPC logging middleware.
func HTTPServer(logger log.Logger) middleware.Middleware {
	log := log.NewHelper("http", logger)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var (
				path   string
				method string
			)
			info, ok := http.FromContext(ctx)
			if ok {
				path = info.Request.RequestURI
				method = info.Request.Method
			}
			reply, err := handler(ctx, req)
			if err != nil {
				log.Errorw(
					"kind", "server",
					"http.path", path,
					"http.method", method,
					"http.code", errors.Code(err),
					"http.error", err.Error(),
				)
				return nil, err
			}
			log.Infow(
				"kind", "server",
				"http.path", path,
				"http.method", method,
				"http.code", 0,
			)
			return reply, nil
		}
	}
}
