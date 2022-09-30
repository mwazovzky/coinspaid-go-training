package lesson2

import "fmt"

func detectSliceResize(max int64) {
	var i int64
	s := make([]int, 1)

	for i = 0; i < max; i++ {
		before := cap(s)
		s = append(s, 1)
		after := cap(s)

		if before != after {
			fmt.Println(before, after, after-before, (after-before)*100/before)
		}
	}
}
