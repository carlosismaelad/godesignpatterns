package main

import (
	"fmt"

	"carlosismaelad.com/godesignpatterns/generator/generator"
)

func main() {
	logs := generator.LogGenerator()

	for log := range logs {
			fmt.Printf("[%s] %s — %s\n", log.Level, log.Time.Format("15:04:05"), log.Message)
	}
}
