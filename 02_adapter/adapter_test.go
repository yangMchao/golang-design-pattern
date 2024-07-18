package adapter

import (
	"log"
	"testing"
)

func TestAdapter(t *testing.T) {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	res := target.Request()
	log.Println(res)
}
