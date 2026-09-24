package main

import "fmt"

// переменная, объявленная внутри функции, видна только в ней.
// переменная на уровне пакета (вне функций) видна всем функциям файла/пакета

var counter int // package-level переменная

func increment() {
	counter++ // видит и меняет package-level переменную напрямую
}

func setValue() {
	secret := 42 // видна только внутри setValue
	fmt.Println("внутри setValue:", secret)
}

func main() {
	increment()
	increment()
	fmt.Println("counter:", counter)

	setValue()

	// поломка: переменная, объявленная внутри другой функции, снаружи не видна
	// fmt.Println(secret)
	// ошибка: undefined: secret

	// у блоков if/for тоже своя область видимости
	if x := 10; x > 5 {
		fmt.Println("x внутри if:", x)
	}
	// fmt.Println(x) // ошибка: undefined: x, переменная умерла вместе с if
}
