package lesson4

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Channels() {
	// basic()
	// blockingRead()
	// nonBlockingRead()
	// timeout()
	channelRange()
}

func basic() {
	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	fmt.Println(<-ch)
}

func blockingRead() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go sendAfter(ch1, 3, 1)
	go sendAfter(ch2, 2, 2)

	select {
	case <-ch1:
		fmt.Println("channel 1")
	case <-ch2:
		fmt.Println("channel 2")
	}
}

func sendAfter(ch chan<- int, timeout int, value int) {
	time.Sleep(time.Second * time.Duration(timeout))
	ch <- value
}

func nonBlockingRead() {
	ch3 := make(chan int)

	go sendAfterRandomTimeout(ch3, 3, 1)
	time.Sleep(time.Second)

	select {
	case <-ch3:
		fmt.Println("channel 3")
	default:
		fmt.Println("default") // we get here if channel is still empty
	}
}

func sendAfterRandomTimeout(ch chan<- int, max int, value int) {
	timeout := rand.Intn(max)
	fmt.Println("timeout: ", timeout)
	time.Sleep(time.Second * time.Duration(timeout))
	ch <- value
}

func timeout() {
	ch := make(chan int)

	go sendAfterRandomTimeout(ch, 3, 42)

	select {
	case value := <-ch:
		fmt.Println("channel", value)
	case t := <-time.After(time.Second):
		fmt.Println("timeout", t)
	}
}

func channelRange() {
	mshCg := make(chan string)
	done := make(chan struct{})

	go func() {
		t := time.NewTicker(time.Second)
		defer t.Stop()

		for {
			select {
			case <-t.C:
			case <-done:
				return
			}
			mshCg <- "MSG1"
		}
	}()

	go func() {
		t := time.NewTicker(time.Millisecond * 1500)
		defer t.Stop()

		for {
			select {
			case <-t.C:
			case <-done:
				return
			}
			mshCg <- "MSG2"
		}
	}()

	go func() {
		time.Sleep(time.Second * 5)
		close(done)
		close(mshCg)
	}()

	for msg := range mshCg {
		fmt.Println(msg)
	}
}
