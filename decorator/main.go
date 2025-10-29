package main

import (
	"fmt"

	"carlosismaelad.com/godesignpatterns/decorator/decorator"
	"carlosismaelad.com/godesignpatterns/decorator/model"
	"carlosismaelad.com/godesignpatterns/decorator/operations"
)

func main() {
	decoratedAdd := decorator.TimeDecorator(operations.Add)
	result := decoratedAdd(model.Input{A: 20, B: 50})
	fmt.Println(result)
}