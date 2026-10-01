package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // ГОНКА ДАННЫХ! Это на самом деле 3 действия: прочитать, +1, записать
		}()
	}
	wg.Wait()
	fmt.Println("Ожидали 1000, получили:", counter) // 987, 993, 1000... каждый раз по-разному

	// Запусти несколько раз - числа будут разные. Очень наглядно.
	// Затем в терминале запусти детектор гонок:
	//   go run -race main.go
	// Он покажет: WARNING: DATA RACE
	// Запомните флаг -race, он спасёт вам не одну ночь.
}
