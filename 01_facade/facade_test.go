package facade

import (
	"log"
	"testing"
)

// TestFacadeAPI ...
func TestFacadeAPI(t *testing.T) {
	api := NewAPI()
	ret := api.Test()
	log.Println(ret)
}
