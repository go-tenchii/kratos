package main

import (
	"context"
	"log"

	"github.com/go-kratos/kratos/v2/errors"
	pb "github.com/go-kratos/kratos/v2/examples/helloworld/helloworld"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/status"
	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
)

func main() {
	callHTTP()
	callGRPC()
}

func callHTTP() {
	client, err := transhttp.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.Get("http://127.0.0.1:8000/helloworld/kratos")
	if err != nil {
		log.Fatal(err)
	}
	if err := transhttp.CheckResponse(res); err != nil {
		log.Fatal(err)
	}
	reply := new(pb.HelloReply)
	if err := transhttp.DecodeResponse(res, &reply); err != nil {
		log.Fatal(err)
	}
	log.Printf("[http] hello %s\n", reply.Message)

	// returns error
	res, err = client.Get("http://127.0.0.1:8000/helloworld/error")
	if err != nil {
		log.Fatal(err)
	}
	if err = transhttp.CheckResponse(res); err != nil {
		log.Printf("[http] SayHello error: %v\n", err)
	}
	if errors.IsInvalidArgument(err) {
		log.Printf("[http] SayHello error is invalid argument: %v\n", err)
	}
}

func callGRPC() {
	conn, err := transgrpc.NewClient(
		"127.0.0.1:9000",
		transgrpc.WithInsecure(),
		transgrpc.WithMiddleware(
			middleware.Chain(
				status.Client(),
				recovery.Recovery(),
			),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewGreeterClient(conn)
	reply, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "kratos"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[grpc] SayHello %+v\n", reply)

	// returns error
	_, err = client.SayHello(context.Background(), &pb.HelloRequest{Name: "error"})
	if err != nil {
		log.Printf("[grpc] SayHello error: %v\n", err)
	}
	if errors.IsInvalidArgument(err) {
		log.Printf("[grpc] SayHello error is invalid argument: %v\n", err)
	}
}
