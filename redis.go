package redi2go

import (
	"time"

	"github.com/ahmadmuzakki29/redigomock"
	redigo "github.com/garyburd/redigo/redis"
)

type Redis struct {
	pool redigo.Pool
	Mock *redigomock.Conn
}

type Reply struct {
	value interface{}
	err   *Error
}

func GetRedisMock() (*Redis, *redigomock.Conn) {
	mock := redigomock.NewConn()

	return &Redis{Mock: mock}, mock
}

func GetRedis(address string) *Redis {
	pool := redigo.Pool{
		MaxIdle:   10,
		MaxActive: 16,
		Dial: func() (redigo.Conn, error) {
			return redigo.Dial(
				"tcp", address,
				redigo.DialConnectTimeout(2*time.Second),
			)
		},
	}

	return &Redis{pool, nil}
}

type Error struct {
	error
}

func (e *Error) ErrNil() bool {
	return e.error == redigo.ErrNil
}
