package main

import "fmt"

// Урок 8: "сломай специально" - типичные ошибки новичков с указателями

func addOne(n *int) {
	// n = n + 1 // поломка 3: invalid operation (mismatched types *int and int)
	*n = *n + 1
}

func main() {
	// поломка 1 - разыменование nil-указателя
	var p *int
	// fmt.Println(*p)
	// panic: runtime error: invalid memory address or nil pointer dereference
	_ = p

	// поломка 2 - забыли & при вызове функции с параметром-указателем
	x := 5
	// addOne(x) // cannot use x (variable of type int) as *int value
	addOne(&x)
	fmt.Println("x =", x) // 6

	// поломка 4 - в Go нет арифметики указателей (в отличие от C)
	q := &x
	// q++ // invalid operation: q++ (non-numeric type *int)
	fmt.Println("q =", q)
}

// скажи: "эту панику (поломка 1) вы увидите ещё много раз в карьере.
// теперь вы знаете, что она значит"
