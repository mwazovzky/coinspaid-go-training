package lesson4

import (
	"fmt"
	"time"
)

func Homework() {
	// m := map[int]int{}
	// for i := 0; i < 100; i++ {
	// 	go func(ii int) {
	// 		m[ii] = ii
	// 	}(i)
	// }
	// panic

	sm := NewSyncMap()
	for i := 0; i < 100; i++ {
		go func(ii int) {
			sm.Set(ii, ii)
		}(i)
	}

	time.Sleep(time.Second)

	fmt.Println(sm.data)
}
