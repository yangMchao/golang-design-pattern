package factorymethod

import (
	"log"
	"testing"
)

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperator(t *testing.T) {
	var (
		factory OperatorFactory
	)
	factory = PlusOperatorFactory{}
	computePlus := compute(factory, 1, 2)
	log.Printf("plus computePlus: %d \n", computePlus)
	factory = &MinusOperatorFactory{}
	computeMinus := compute(factory, 1, 2)
	log.Printf("minus computeMinus: %d \n", computeMinus)
}
