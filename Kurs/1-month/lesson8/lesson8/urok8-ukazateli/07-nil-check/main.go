package main

import "fmt"

// Урок 8, Шаг 7: nil - пустой указатель, и проверка перед разыменованием

func main() {
	var p *int
	fmt.Println(p == nil) // true
	fmt.Println(p)        // <nil>

	// fmt.Println(*p)
	// panic: runtime error: invalid memory address or nil pointer dereference

	if p != nil {
		fmt.Println(*p)
	} else {
		fmt.Println("Указатель пустой")
	}

	x := 7
	p = &x
	if p != nil {
		fmt.Println("Теперь есть:", *p)
	}
}

// раскомментируй строку с паникой, покажи ошибку в консоли, потом верни обратно
