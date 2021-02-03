package mongodb

import (
	"github.com/go-kratos/kratos/pkg/net/netutil/breaker"
	"github.com/go-kratos/kratos/pkg/time"
)

type Config struct {
	DSN          string          // write data source name.
	SocketTimeout time.Duration
	PoolLimit 	  int
	Breaker      *breaker.Config // breaker
}
