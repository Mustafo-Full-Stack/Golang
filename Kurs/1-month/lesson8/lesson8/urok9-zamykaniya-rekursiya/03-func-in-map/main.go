package main

import "fmt"

// Урок 9, Шаг 3: функции в коллекциях - таблица операций вместо switch

func main() {
	operations := map[string]func(int, int) int{
		"add": func(a, b int) int { return a + b },
		"sub": func(a, b int) int { return a - b },
		"mul": func(a, b int) int { return a * b },
	}

	for name, op := range operations {
		fmt.Printf("%s(6, 3) = %d\n", name, op(6, 3))
	}
}

// красивая замена switch - таблица операций, свяжи с калькулятором из урока 7
