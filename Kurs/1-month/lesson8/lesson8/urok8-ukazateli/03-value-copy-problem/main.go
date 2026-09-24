package main

import "fmt"

// Урок 8, Шаг 3: проблема - функция получает копию аргумента,
// изменение копии не видно снаружи

func addOne(n *int) {
	*n = *n + 1
	fmt.Println("Внутри функции n =", *n)
}

func main() {
	x := 5
	addOne(&x)
	fmt.Println("Снаружи x =", x) // 5
}

