package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// Демонстрация утечки горутин - зачем на самом деле нужен context.

func leaky() {
	ch := make(chan int)
	go func() {
		val := <-ch // ждёт вечно
		fmt.Println(val)
	}()
	// функция завершилась, горутина висит навсегда
}

func fixed(ctx context.Context) {
	ch := make(chan int)
	go func() {
		select {
		case val := <-ch:
			fmt.Println(val)
		case <-ctx.Done(): // умеем остановиться по сигналу
			return
		}
	}()
}

func main() {
	fmt.Println("Горутин до:", runtime.NumGoroutine())
	for i := 0; i < 100; i++ {
		leaky()
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Горутин после leaky:", runtime.NumGoroutine()) // 101!

	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 100; i++ {
		fixed(ctx)
	}
	cancel() // отменяем разом - все горутины из fixed завершатся
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Горутин после fixed+cancel:", runtime.NumGoroutine())
}
