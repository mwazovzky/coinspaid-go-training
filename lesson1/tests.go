package lesson1

import "fmt"

func Tests() {
	// Exercise 1
	fmt.Println("Числа фиббоначи")
	next := Next()
	for i := 0; i < 10; i++ {
		fmt.Println(next())
	}

	// Exercise 2
	fmt.Println("Арифметическая прогрессия")
	rule := func(value int) int {
		diff := 2
		return value + diff
	}
	generator := Generator(rule)
	for i := 0; i < 10; i++ {
		fmt.Println(generator())
	}

	// Exercise 3
	fmt.Println("Проверка строки на наличие повторяющихся символов")
	fmt.Println(CheckDuplicates("abc"))  // false
	fmt.Println(CheckDuplicates("aa"))   // true
	fmt.Println(CheckDuplicates("Aa"))   // true
	fmt.Println(CheckDuplicates("aA"))   // true
	fmt.Println(CheckDuplicates("abca")) // true
	fmt.Println(CheckDuplicates("abcA")) // true

	// Exercise 4
	fmt.Println("Калькулятор")
	v1 := Eval(10, 20, Plus, "45", Minus)
	fmt.Printf("%T %v\n", v1, v1) // float64 -15
	v2 := Eval(10, 2.5, Plus)
	fmt.Printf("%T %v\n", v2, v2) // float64 12.5
}
