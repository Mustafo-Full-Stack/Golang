package main

import "fmt"

// Урок 8, Шаг 5: классика - swap

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	x, y := 1, 2
	fmt.Println("До: ", x, y)
	swap(&x, &y)
	fmt.Println("После:", x, y)
}

// func copySwap(a, b int) {
// 	a, b = b, a
// }

// func main() {
// 	x, y := 1, 2
// 	fmt.Println("До: ", x, y)
// 	copySwap(x, y)
// 	fmt.Println("После:", x, y)
// }
