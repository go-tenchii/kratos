package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/errors"
	pb "github.com/go-kratos/kratos/v2/examples/helloworld/helloworld"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/log/stdlog"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/status"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	if in.Name == "error" {
		return nil, errors.InvalidArgument("BadRequest", "invalid argument %s", in.Name)
	}
	if in.Name == "panic" {
		panic("grpc panic")
	}
	return &pb.HelloReply{Message: fmt.Sprintf("Hello %+v", in)}, nil
}

func loggerInfo(logger log.Logger) middleware.Middleware {
	log := log.NewHelper("logger2", logger)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {

			tr, ok := transport.FromContext(ctx)
			if ok {
				log.Infof("transport: %+v", tr)
			}
			h, ok := http.FromContext(ctx)
			if ok {
				log.Infof("http: [%s] %s", h.Request.Method, h.Request.URL.Path)
			}
			g, ok := grpc.FromContext(ctx)
			if ok {
				log.Infof("grpc: %s", g.FullMethod)
			}

			return handler(ctx, req)
		}
	}
}

func main() {
	logger := stdlog.NewLogger(stdlog.Writer(os.Stdout))
	defer logger.Close()

	log := log.NewHelper("main", logger)

	s := &server{}
	app := kratos.New()

	httpSrv := http.NewServer(
		http.Address(":8000"),
		http.Middleware(
			middleware.Chain(
				logging.HTTPServer(logger),
				recovery.Recovery(),
			),
		))
	grpcSrv := grpc.NewServer(
		grpc.Address(":9000"),
		grpc.Middleware(
			middleware.Chain(
				logging.GRPCServer(logger),
				status.Server(),
				recovery.Recovery(),
			),
		))

	pb.RegisterGreeterServer(grpcSrv, s)
	pb.RegisterGreeterHTTPServer(httpSrv, s)

	app.Append(httpSrv)
	app.Append(grpcSrv)

	if err := app.Run(); err != nil {
		log.Error(err)
	}
}
