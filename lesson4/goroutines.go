package lesson4

import (
	"fmt"
	"sync"
	"time"
)

func Goroutines() {

	// runtime.GOMAXPROCS(1)

	counter := 0
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second / 5)
	fmt.Println(counter)

	counter = 0
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			counter++
		}()
	}
	time.Sleep(time.Second / 5)
	fmt.Println(counter)

	var murw sync.RWMutex
	for i := 0; i <= 3; i++ {
		go func() {
			murw.RLock()
			time.Sleep(time.Second)
			fmt.Println("RLock", i)
			murw.RUnlock()
		}()
	}
	time.Sleep(time.Second * 2)

	for i := 0; i <= 3; i++ {
		go func() {
			murw.Lock()
			time.Sleep(time.Second / 4)
			fmt.Println("Lock", i)
			murw.Unlock()
		}()
	}
	time.Sleep(time.Second * 2)

	wg := sync.WaitGroup{}
	wg.Add(4)
	for i := 0; i <= 3; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
