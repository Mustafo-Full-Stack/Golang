package main

import (
	"context"
	"fmt"
	"time"
)

func workerCtx(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер %d: остановлен (%v)\n", id, ctx.Err())
			return
		default:
			fmt.Printf("Воркер %d: работаю\n", id)
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 1; i <= 3; i++ {
		go workerCtx(ctx, i)
	}

	time.Sleep(1 * time.Second)
	cancel() // отменяем всё дерево
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Готово")

	// Это стандарт индустрии. В любой Go-библиотеке первым параметром идёт ctx.
}
