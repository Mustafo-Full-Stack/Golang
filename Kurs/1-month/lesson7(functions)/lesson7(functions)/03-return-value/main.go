package main

import "fmt"

// функция с возвратом: после параметров указываем тип результата,
// а в теле обязательно есть return с значением этого типа

func celsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func main() {
	for c := 0.0; c <= 100; c += 20 {
		fmt.Printf("%.0f°C = %.1f°F\n", c, celsiusToFahrenheit(c))
	}

	// результат функции можно сразу передать в другую функцию
	fmt.Println("100°C в фаренгейтах:", celsiusToFahrenheit(100))
}
