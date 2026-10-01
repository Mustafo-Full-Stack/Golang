package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("Ядер CPU:      ", runtime.NumCPU())
	fmt.Println("GOMAXPROCS:    ", runtime.GOMAXPROCS(0))
	fmt.Println("Горутин сейчас:", runtime.NumGoroutine())

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
		}()
	}
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Во время работы:", runtime.NumGoroutine())
	wg.Wait()
}
