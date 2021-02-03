package mongodb

import (
	"github.com/go-tenchii/kratos/pkg/net/netutil/breaker"
	"github.com/go-tenchii/kratos/pkg/time"
)

type Config struct {
	DSN          string          // write data source name.
	SocketTimeout time.Duration
	PoolLimit 	  int
	Breaker      *breaker.Config // breaker
}
