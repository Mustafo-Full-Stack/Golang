package main

import "fmt"

// Урок 8, Шаг 10: указатель на массив vs срез

func modifyArray(arr [3]int) {
	arr[0] = 999 // копия!
}

func modifyArrayPtr(arr *[3]int) {
	arr[0] = 999 // Go сам разыменует
}

func modifySlice(s []int) {
	s[0] = 999 // работает без указателя
}

func main() {
	a := [3]int{1, 2, 3}
	modifyArray(a)
	fmt.Println("Массив по значению:", a) // [1 2 3]

	modifyArrayPtr(&a)
	fmt.Println("Массив по указателю:", a) // [999 2 3]

	s := []int{1, 2, 3}
	modifySlice(s)
	fmt.Println("Срез:", s) // [999 2 3]
}

// ключевой вывод: срезам и map указатели не нужны - они уже ссылочные.
// это снимает частый вопрос студентов
