package main

import (
	"fmt"
	"time"
)

// Урок 9, Шаг 10: рекурсия - Фибоначчи и её цена,
// решение - мемоизация через замыкание

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func makeFib() func(int) int {
	cache := map[int]int{}
	var fib func(int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		if v, ok := cache[n]; ok {
			return v
		}
		res := fib(n-1) + fib(n-2)
		cache[n] = res
		return res
	}
	return fib
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(fib(i), " ") // 0 1 1 2 3 5 8 13 21 34
	}
	fmt.Println()

	// покажи проблему - наивный fib делает миллионы повторных вычислений.
	// на современном железе fib(35) уже не секунды, а миллисекунды,
	// поэтому для эффекта "несколько секунд" бери fib(45)
	start := time.Now()
	fmt.Println(fib(45))
	fmt.Println("Заняло:", time.Since(start)) // на слабой машине - несколько секунд, на быстрой - пара секунд

	// решение - замыкание + рекурсия + map
	fastFib := makeFib()
	start = time.Now()
	fmt.Println(fastFib(80))
	fmt.Println("Заняло:", time.Since(start)) // микросекунды
}

// кульминация урока: эффект от сравнения времени очень сильный
