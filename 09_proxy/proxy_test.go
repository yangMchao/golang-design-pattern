package proxy

import (
	"log"
	"testing"
)

func TestProxy(t *testing.T) {
	var sub Subject
	sub = &Proxy{}
	res := sub.Do()
	log.Println(res)
}
