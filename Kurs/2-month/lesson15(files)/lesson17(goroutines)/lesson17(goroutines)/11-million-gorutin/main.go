package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	start := time.Now()

	const n = 1_000_000
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond)
		}()
	}
	wg.Wait()
	fmt.Printf("Запустили и дождались %d горутин за %v\n", n, time.Since(start))

	// Впечатляющий момент курса: миллион "потоков" на обычном ноутбуке
	// за пару секунд. Попробуйте создать миллион потоков ОС - компьютер умрёт.
}
