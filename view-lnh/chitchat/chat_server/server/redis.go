package server

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"time"
)

func InitRedis(addr string, idleConn, maxConn int, idleTimeout time.Duration) (pool *redis.Pool, err error) {
	defer func() {
		if ferr := recover(); ferr != nil {
			err = errors.New("Init redis err " + ferr.(string))
		}
	}()
	pool = &redis.Pool{
		MaxIdle:     idleConn,
		MaxActive:   maxConn,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}
	return
}
