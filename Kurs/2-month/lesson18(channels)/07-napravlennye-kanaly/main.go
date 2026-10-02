package main

import "fmt"

func generate(out chan<- int, n int) { // только пишет
	for i := 1; i <= n; i++ {
		out <- i * i
	}
	close(out)
}

func print(in <-chan int, done chan<- bool) { // только читает
	for v := range in {
		fmt.Print(v, " ")
	}
	fmt.Println()
	done <- true
}

func main() {
	numbers := make(chan int)
	done := make(chan bool)

	go generate(numbers, 10)
	go print(numbers, done)

	<-done
	fmt.Println("Готово")

	// Компилятор проверит направление сам. Попробуй в generate написать
	// v := <-out - будет ошибка компиляции:
	// invalid operation: cannot receive from send-only channel
}
