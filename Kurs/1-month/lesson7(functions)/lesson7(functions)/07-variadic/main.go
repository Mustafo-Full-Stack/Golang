package main

import "fmt"

// "..." перед типом означает "сколько угодно аргументов".
// внутри функции параметр - обычный срез. именно так устроен fmt.Println

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// вариативный параметр может быть только последним,
// но перед ним могут идти обычные обязательные параметры
func greetAll(greeting string, names ...string) {
	for _, n := range names {
		fmt.Printf("%s, %s!\n", greeting, n)
	}
}

func main() {
	fmt.Println(sum())              // 0
	fmt.Println(sum(1, 2, 3))       // 6
	fmt.Println(sum(1, 2, 3, 4, 5)) // 15

	// можно передать готовый срез - снова три точки, как у append в уроке про срезы
	nums := []int{10, 20, 30}
	fmt.Println(sum(nums...)) // 60

	greetAll("Привет", "Али", "Вера", "Тимур")

	// поломка: вариативный параметр можно ставить только последним
	// func bad(nums ...int, name string) {}
	// ошибка: can only use ... with final parameter in list
}
