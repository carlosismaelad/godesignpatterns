package operations

import (
	"time"

	"carlosismaelad.com/godesignpatterns/decorator/model"
)

func Add(input model.Input) int {
	time.Sleep(3 * time.Second)
	return input.A + input.B
}