package main

import (
	"fmt"

	"carlosismaelad.com/godesignpatterns/builder/builder"
)


func main() {
	c := builder.NewBuilder().
		Name("Carlos Dourado").
		Email("carlos@example.com").
		Phone("85987654321").
		Address("123 Main St").
		Age(33).
		IsActive(true).
		Build()

	fmt.Printf("%+v\n", c)
}