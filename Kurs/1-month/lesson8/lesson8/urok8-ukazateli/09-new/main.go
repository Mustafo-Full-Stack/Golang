package main

import "fmt"

// Урок 8, Шаг 9: new(T) - создаёт нулевое значение типа T и возвращает указатель на него

func main() {
	// p := new(int)
	// // var p *int
	// fmt.Println(p, *p) // адрес и 0
	// *p = 42
	// fmt.Println(*p) // 42

	// эквивалентно и понятнее - так пишут на практике чаще:
	var x int // x :=0
	q := &x
	*q = 42
	fmt.Println(*q)
}
