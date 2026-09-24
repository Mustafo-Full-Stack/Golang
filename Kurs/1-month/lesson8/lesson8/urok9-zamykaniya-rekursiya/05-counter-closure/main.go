package main

import "fmt"

// Урок 9, Шаг 5: замыкание со состоянием - счётчик.
// центральный момент урока

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	next := counter()
	fmt.Println(next()) // 1
	fmt.Println(next()) // 2
	fmt.Println(next()) // 3

	other := counter()
	fmt.Println(other()) // 1 - независимый счётчик!
	fmt.Println(next())  // 4 - первый продолжает
}

// проговори медленно: count объявлена внутри counter, функция counter уже
// завершилась, но переменная жива, потому что внутренняя функция её держит.
// каждый вызов counter() создаёт новую переменную count
