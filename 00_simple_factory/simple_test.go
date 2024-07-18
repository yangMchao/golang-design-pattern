package simplefactory

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

// TestType1 test get hi api with factory
func TestType1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Tom")
	typeOf := reflect.TypeOf(api)
	if s != "Hi, Tom" {
		t.Fatal("Type1 test fail")
	}
	log.Println(fmt.Sprintf("type %s , value %s", typeOf, s))
}

func TestType2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("Tom")
	typeOf := reflect.TypeOf(api)
	if s != "Hi, Tom" {
		t.Fatal("Type1 test fail")
	}
	log.Println(fmt.Sprintf("type %s , value %s", typeOf, s))
}
