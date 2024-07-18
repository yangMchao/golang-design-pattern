package prototype

import (
	"log"
	"testing"
)

func init() {
	manager = NewPrototypeManager()

	t1 := &Type1{
		name: "type1",
		D: &DCopy{
			dDCopy: "d",
		},
	}
	manager.Set("t1", t1)
}

var manager *PrototypeManager

type DCopy struct {
	dDCopy string
}

type Type1 struct {
	name string
	D    *DCopy
}

func (t *Type1) Clone() Cloneable {
	tc := *t
	return &tc
}

type Type2 struct {
	name string
}

func (t *Type2) Clone() Cloneable {
	tc := *t
	return &tc
}

func TestCopy(t *testing.T) {
	dCopy := &DCopy{
		dDCopy: "d",
	}
	t1 := &Type1{
		name: "type1",
		D:    dCopy,
	}
	tc := *t1
	t1.D.dDCopy = "c"
	log.Println(tc.D.dDCopy)
	log.Println(t1.D.dDCopy)
}

func TestClone(t *testing.T) {
	t1 := manager.Get("t1")

	t2 := t1.Clone()

	if t1 == t2 {
		t.Fatal("error! get clone not working")
	}
}

func TestCloneFromManager(t *testing.T) {
	c := manager.Get("t1").Clone()
	t1 := c.(*Type1)
	t2 := c.(*Type1)
	if t1.name != "type1" {
		t.Fatal("error")
	}
	t1.D.dDCopy = "c"
	log.Println(t2.D.dDCopy)
	log.Println(t1.D.dDCopy)
}
