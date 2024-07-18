package builder

import (
	"log"
	"testing"
)

// go test -v -run=TestBuilder1
// go test -v -run=TestBuilder2

func TestBuilder1(t *testing.T) {
	builder := &Builder1{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	log.Printf("stringBuilder  ：%s \n", res)

}

func TestBuilder2(t *testing.T) {
	builder := &Builder2{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	log.Printf("intBuillder：%d \n", res)

}
