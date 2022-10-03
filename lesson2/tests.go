package lesson2

import "fmt"

func Tests() {
	// Exercise 1
	fmt.Println("Определить как меняется размер слайса при добавлении элементов")
	detectSliceResize(1000000)
}
