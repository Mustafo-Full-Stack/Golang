package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Print(n, " ")
		}(i)
	}
	wg.Wait()
	fmt.Println()

	// Запусти несколько раз подряд (go run main.go) - порядок чисел будет разным.
	// Вывод: никогда не полагайтесь на порядок выполнения горутин.
}
