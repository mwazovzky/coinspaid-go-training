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

	// Exercise 3
	fmt.Println("Проверка строки на наличие повторяющихся символов")
	fmt.Println(lesson1.CheckDuplicates("abc"))  // false
	fmt.Println(lesson1.CheckDuplicates("aa"))   // true
	fmt.Println(lesson1.CheckDuplicates("Aa"))   // true
	fmt.Println(lesson1.CheckDuplicates("aA"))   // true
	fmt.Println(lesson1.CheckDuplicates("abca")) // true
	fmt.Println(lesson1.CheckDuplicates("abcA")) // true

	// Exercise 4
	fmt.Println("Калькулятор")
	v1 := lesson1.Eval(10, 20, lesson1.Plus, "45", lesson1.Minus)
	fmt.Printf("%T %v\n", v1, v1) // float64 -15
	v2 := lesson1.Eval(10, 2.5, lesson1.Plus)
	fmt.Printf("%T %v\n", v2, v2) // float64 12.5
}
