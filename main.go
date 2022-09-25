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
}
