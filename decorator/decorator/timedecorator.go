package decorator

import (
	"fmt"
	"time"

	"carlosismaelad.com/godesignpatterns/decorator/model"
)

func TimeDecorator(fn func(model.Input) int) func(model.Input) int {
	return func(input model.Input) int {
		start := time.Now()
		result := fn(input)
		duration := time.Since(start)
		fmt.Printf("Tempo de execução: %dms\n\n", duration.Milliseconds())
		return result
	}
}