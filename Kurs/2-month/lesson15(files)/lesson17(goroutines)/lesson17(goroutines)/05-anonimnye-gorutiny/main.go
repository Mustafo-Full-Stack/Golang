package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) { // передаём i как аргумент - ловушка замыкания из среза/слайсов!
			defer wg.Done()
			fmt.Printf("Горутина %d работает\n", id)
			time.Sleep(time.Duration(id) * 100 * time.Millisecond)
			fmt.Printf("Горутина %d завершена\n", id)
		}(i)
	}
	wg.Wait()

	// Почему id передаём аргументом, а не берём i из внешней области?
	// Так безопаснее и понятнее в любой версии Go (с Go 1.22 цикл и так
	// создаёт свою переменную на каждой итерации, но привычка передавать
	// аргументом остаётся хорошим стилем).
}
