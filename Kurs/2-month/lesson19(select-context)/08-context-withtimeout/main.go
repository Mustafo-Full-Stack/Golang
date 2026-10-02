package main

import (
	"context"
	"fmt"
	"time"
)

func fetchData(ctx context.Context, delay time.Duration) (string, error) {
	select {
	case <-time.After(delay):
		return "данные получены", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func main() {
	// успеваем
	ctx1, cancel1 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel1()

	data, err := fetchData(ctx1, 300*time.Millisecond)
	fmt.Println(data, err)

	// не успеваем
	ctx2, cancel2 := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel2()

	data, err = fetchData(ctx2, 2*time.Second)
	fmt.Println(data, err) // "" context deadline exceeded

	// defer cancel() обязателен - освобождает ресурсы, даже если таймаут не сработал.
}
