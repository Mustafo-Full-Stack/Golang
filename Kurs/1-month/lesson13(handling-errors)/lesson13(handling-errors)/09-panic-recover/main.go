package main

import "fmt"

func mustPositive(n int) int {
	if n <= 0 {
		panic(fmt.Sprintf("ожидалось положительное число, получено %d", n))
	}
	return n
}

func safeCall() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Перехватили панику:", r)
		}
	}()

	fmt.Println(mustPositive(5))
	fmt.Println(mustPositive(-1)) // паника
	fmt.Println("Эта строка не выполнится")
}

func main() {
	safeCall()
	fmt.Println("Программа продолжает работать")

	// Правило: panic - для ошибок программиста (нарушенный инвариант),
	// error - для ожидаемых ситуаций (неверный ввод, файла нет, сети нет).
	//
	// ПЛОХО:
	// func getUser(id int) string {
	//     if id <= 0 {
	//         panic("bad id") // нет! это обычная ошибка
	//     }
	//     ...
	// }
	//
	// ХОРОШО:
	// func getUser(id int) (string, error) {
	//     if id <= 0 {
	//         return "", ErrInvalidInput
	//     }
	//     ...
	// }
}
