package main

import "fmt"

// Урок 9, Шаг 4: функция возвращает функцию

func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func main() {
	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println(double(5)) // 10
	fmt.Println(triple(5)) // 15
}

// это уже замыкание: возвращённая функция помнит свой factor
