package main

import "fmt"

// если подряд идут параметры одного типа - тип пишут один раз:
// func max(a, b int) int   вместо   func max(a int, b int) int

func max(a, b int) int {
	if a > b {
		return a // как только знаем ответ - сразу выходим из функции
	}
	return b
}

func classify(age int) string {
	if age < 0 {
		return "некорректный возраст"
	}
	if age < 18 {
		return "несовершеннолетний"
	}
	if age < 65 {
		return "взрослый"
	}
	return "пенсионер"
}

func main() {
	fmt.Println(max(10, 7))
	fmt.Println(classify(25))
	fmt.Println(classify(-5))

	// ранний return чище, чем городить if/else на каждый случай:
	// как только знаем ответ - выходим

	// поломка: если у функции есть тип результата, она обязана
	// вернуть значение на любом пути выполнения
	// func square(x int) int {
	// 	fmt.Println(x * x)
	// }
	// ошибка компиляции: missing return
}
