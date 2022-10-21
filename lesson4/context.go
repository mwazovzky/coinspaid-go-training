package lesson4

import (
	"context"
	"fmt"
	"time"
)

type ValueKeyType string

func Context() {
	contextWithValue()
	contextWithTimeout()
	contextWithCancel()
}

func contextWithValue() {
	var ctx context.Context

	ctx = context.Background() // empty context
	ctx = context.TODO()       // indication that context is required

	key := ValueKeyType("key")
	ctxWithValue := context.WithValue(ctx, key, "value")
	value := ctxWithValue.Value(key)
	fmt.Println(key, value)
}

func contextWithTimeout() {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := asyncFunc(ctxWithTimeout)
	fmt.Println(res, err)

	time.Sleep(3 * time.Second)
}

func asyncFunc(ctx context.Context) (string, error) {
	select {
	case <-time.After(time.Second * 2):
		fmt.Println("data loaded")
		return "data", nil
	case <-ctx.Done():
		fmt.Println("data load failed")
		return "", fmt.Errorf("failed to load data: %w", ctx.Err())
	}
}

func contextWithCancel() {
	ctxWithCancel, cancelFunc := context.WithCancel(context.Background())

	go worker(ctxWithCancel)

	go func() {
		time.Sleep(3 * time.Second)
		cancelFunc()
	}()

	time.Sleep(5 * time.Second)
}

func worker(ctx context.Context) {
	t := time.NewTicker(time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			fmt.Println("working")
		case <-ctx.Done():
			fmt.Println("gracefully stopped")
			return
		}
	}
}
