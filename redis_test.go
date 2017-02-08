package redi2go

import (
	"fmt"
	"testing"
)

func TestReply_String(t *testing.T) {
	redis := GetRedis("devel-redis.tkpd:6379")
	ok, err := redis.Set("jeki", "hello").String()
	fmt.Println(ok, err)
}

func TestErrNil(t *testing.T) {
	redis := GetRedis("devel-redis.tkpd:6379")
	_, err := redis.Get("asdfasdf").String()

	if err.ErrNil() {
		fmt.Println("this is nil")
	}
}
