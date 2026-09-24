package main

import "fmt"

// Урок 8, Шаг 6: практический пример - функции-модификаторы,
// которые меняют значение по адресу

func applyDiscount(price *float64, percent float64) {
	*price = *price * (1 - percent/100)
}

func resetCounter(c *int) {
	*c = 0
}

func main() {
	price := 1000.0
	applyDiscount(&price, 20)
	fmt.Printf("Цена со скидкой: %.2f\n", price) // 800.00

	counter := 57
	resetCounter(&counter)
	fmt.Println("Счётчик:", counter) // 0
}
