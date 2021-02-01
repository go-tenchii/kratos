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
		hc httpconf.Server
		gc grpcconf.Server
	)
	if err := conf.Value("http.server").Scan(&hc); err != nil {
		panic(err)
	}
	if err := conf.Value("grpc.server").Scan(&gc); err != nil {
		panic(err)
	}

	if err := conf.Watch("service.name", func(key string, value config.Value) {
		log.Printf("config changed: %s = %v\n", key, value)
	}); err != nil {
		panic(err)
	}

	log.Printf("http: %s\n", hc.String())
	log.Printf("grpc: %s\n", gc.String())

	// http.Apply(hc)
	// grpc.Apply(gc)

	<-make(chan struct{})
}
