package mongodb

import (
	"gopkg.in/mgo.v2"
	"time"
)

func NewMongodb(c *Config) (session *mgo.Session, err error){
	session, err = mgo.Dial(c.DSN)
	if err != nil {
		return
	}
	session.SetSocketTimeout(time.Duration(c.SocketTimeout))
	session.SetPoolLimit(c.PoolLimit)
	return
}
