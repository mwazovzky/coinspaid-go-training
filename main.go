package main

import (
	"coinspaid-go/lesson1"
	"fmt"
)

func main() {
	// Exercise 1
	fmt.Println("Числа фиббоначи")
	next := lesson1.Next()
	for i := 0; i < 10; i++ {
		fmt.Println(next())
	}

	// Exercise 2
	fmt.Println("Арифметическая прогрессия")
	rule := func(value int) int {
		diff := 2
		return value + diff
	}
	generator := lesson1.Generator(rule)
	for i := 0; i < 10; i++ {
		fmt.Println(generator())
	}
}
