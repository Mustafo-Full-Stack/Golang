package main

import "fmt"

// Урок 9: "сломай специально" - типичные ошибки новичков

// поломка 1 - рекурсия без выхода (базового случая):
// func infinite(n int) int {
// 	return infinite(n - 1)
// }
// runtime: goroutine stack exceeds 1000000000-byte limit
// fatal error: stack overflow

// поломка 2 - рекурсивная анонимная функция без предварительного объявления:
//
//	fib := func(n int) int {
//		if n < 2 {
//			return n
//		}
//		return fib(n-1) + fib(n-2) // fib ещё не объявлена!
//	}
//
// undefined: fib
//
// решение - сначала объявить переменную-тип, потом присвоить саму функцию
func fixedAnonymousRecursion() {
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(6)) // 8
}

// поломка 3 - ловушка замыкания в цикле (классика собеседований).
// в Go до версии 1.22 печатала 4 4 4, начиная с 1.22 - 1 2 3
func closureLoopTrap() {
	funcs := []func(){}
	for i := 1; i <= 3; i++ {
		funcs = append(funcs, func() {
			fmt.Println(i)
		})
	}
	for _, f := range funcs {
		f()
	}
}

// безопасный способ - локальная копия, работает одинаково в любой версии Go
func closureLoopFixed() {
	funcs := []func(){}
	for i := 1; i <= 3; i++ {
		i := i // локальная копия
		funcs = append(funcs, func() {
			fmt.Println(i)
		})
	}
	for _, f := range funcs {
		f()
	}
}

// поломка 4 - defer в цикле (только для объяснения, не для запуска):
// for i := 0; i < 1000; i++ {
// 	f, _ := os.Open("file.txt")
// 	defer f.Close() // накопится 1000 отложенных закрытий!
// }
// defer срабатывает в конце функции, а не итерации.
// решение - вынести тело цикла в отдельную функцию

func main() {
	fixedAnonymousRecursion()
	closureLoopTrap()
	closureLoopFixed()
}
