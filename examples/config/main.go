package main

import (
	"flag"
	"log"

	grpcconf "github.com/go-kratos/kratos/v2/api/kratos/config/grpc"
	httpconf "github.com/go-kratos/kratos/v2/api/kratos/config/http"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/source/file"
)

var flagconf string

func init() {
	flag.StringVar(&flagconf, "conf", "config.yaml", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	conf := config.New(config.WithSource(
		file.NewSource(flagconf),
	))
	if err := conf.Load(); err != nil {
		panic(err)
	}

	var (
		hc httpconf.ServerConfig
		gc grpcconf.ServerConfig
	)
	if err := conf.Value("server.http").Scan(&hc); err != nil {
		panic(err)
	}
	if err := conf.Value("server.grpc").Scan(&gc); err != nil {
		panic(err)
	}

	// srvhttp.Apply(hc)
	// srvgrpc.Apply(hc)

	log.Printf("http: %s\n", hc.String())
	log.Printf("grpc: %s\n", gc.String())
}
