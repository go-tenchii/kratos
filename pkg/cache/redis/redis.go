// Copyright 2012 Gary Burd
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package redis

import "C"
import (
	"github.com/go-redsync/redsync/v4"
	redsyncredis "github.com/go-redsync/redsync/v4/redis"
	goredislib "github.com/go-redis/redis/v7"
	"github.com/go-redsync/redsync/v4/redis/goredis/v7"
	"github.com/stvp/tempredis"
	"time"

	xtime "github.com/go-tenchii/kratos/pkg/time"
)

// Error represents an error returned in a command reply.
type Error string

func (err Error) Error() string { return string(err) }

// Config client settings.
type Config struct {
	Name         string // redis name, for trace
	Proto        string
	Addr         string
	Auth         string
	Db           int
	DialTimeout  xtime.Duration
	ReadTimeout  xtime.Duration
	WriteTimeout xtime.Duration
	SlowLog      xtime.Duration
}

type Redis struct {
	Client *goredislib.Client
	Rs *redsync.Redsync

	pool redsyncredis.Pool
	conf *Config
}

func NewRedis(c *Config) *Redis {
	server, err := tempredis.Start(tempredis.Config{})
	if err != nil {
		panic(err)
	}
	defer server.Term()

	client := goredislib.NewClient(&goredislib.Options{
		Network: "unix",
		Addr:    server.Socket(),
		Password: c.Auth,
		DB:c.Db,
		DialTimeout: time.Duration(c.DialTimeout),
		WriteTimeout: time.Duration(c.WriteTimeout),
		ReadTimeout: time.Duration(c.ReadTimeout),
	})
	pool := goredis.NewPool(client)
	return &Redis{
		Client: client,
		pool: pool,
		conf: c,
		Rs:redsync.New(pool),
	}
}
