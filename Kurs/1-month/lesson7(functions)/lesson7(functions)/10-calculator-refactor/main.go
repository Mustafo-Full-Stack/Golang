package main

import "fmt"


func add(a, b float64) float64 { return a + b }
func sub(a, b float64) float64 { return a - b }
func mul(a, b float64) float64 { return a * b }

func div(a, b float64) (float64, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func readInput() (float64, string, float64) {
	var a, b float64
	var op string
	fmt.Print("Первое число: ")
	fmt.Scanln(&a)
	fmt.Print("Операция: ")
	fmt.Scanln(&op)
	fmt.Print("Второе число: ")
	fmt.Scanln(&b)
	return a, op, b
}

func main() {
	a, op, b := readInput()

	switch op {
	case "+":
		fmt.Println("=", add(a, b))
	case "-":
		fmt.Println("=", sub(a, b))
	case "*":
		fmt.Println("=", mul(a, b))
	case "/":
		if res, ok := div(a, b); ok {
			fmt.Println("=", res)
		} else {
			fmt.Println("Деление на ноль!")
		}
	default:
		fmt.Println("Неизвестная операция")
	}
}
