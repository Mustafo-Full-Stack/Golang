package main

import (
	"fmt"
	"sync"
)

// Fan-in - объединение нескольких источников в один поток данных.
func merge(chans ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, c := range chans {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func main() {
	a := gen(1, 2, 3)
	b := gen(4, 5, 6)
	c := gen(7, 8, 9)

	sum := 0
	for v := range merge(a, b, c) {
		sum += v
	}
	fmt.Println("Сумма всех значений из трёх каналов:", sum) // 45
}
