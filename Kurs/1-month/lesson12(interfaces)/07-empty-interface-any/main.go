package main

import "fmt"

type Circle struct{ Radius float64 }

// any - это синоним interface{} начиная с Go 1.18. Он не требует ни одного
// метода, поэтому подходит абсолютно любой тип. Удобно, но проверки типов
// на этапе компиляции пропадают - используем осознанно, а не "на всякий случай".
func printAnything(v any) {
	fmt.Printf("Значение: %v, Тип: %T\n", v, v)
}

// var s interface{} // s - пустой интерфейс



func main() {
	printAnything(42)
	printAnything("привет")
	printAnything(3.14)
	printAnything(true)
	printAnything([]int{1, 2, 3})
	printAnything(Circle{Radius: 5})
	printAnything(map[string]int{"a": 1})
}
