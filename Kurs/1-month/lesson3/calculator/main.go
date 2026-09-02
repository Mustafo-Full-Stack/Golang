package main

import "fmt"

func main() {
	var (
		a, b float64
		op   string
	)

	fmt.Println("=== Калькулятор ===")
	fmt.Print("Первое число: ")
	fmt.Scanln(&a)
	fmt.Print("Операция (+ - * /): ")
	fmt.Scanln(&op)
	fmt.Print("Второе число: ")
	fmt.Scanln(&b)

	switch op {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f\n", a, b, a+b)
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f\n", a, b, a-b)
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f\n", a, b, a*b)
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль")
		} else {
			fmt.Printf("%.2f / %.2f = %.2f\n", a, b, a/b)
		}
	default:
		fmt.Printf("Неизвестная операция: %q\n", op)
	}
}
