package main

import "fmt"

// Pipeline (конвейер) - самый красивый паттерн Go.
// Каждая стадия - своя горутина, данные текут по трубе.
// Похоже на пайпы в Unix: cat file | grep x | sort

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func filterEven(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
	}()
	return out
}

func main() {
	// три стадии, соединённые каналами
	for v := range filterEven(square(generator(1, 2, 3, 4, 5, 6))) {
		fmt.Print(v, " ") // 4 16 36
	}
	fmt.Println()
}
