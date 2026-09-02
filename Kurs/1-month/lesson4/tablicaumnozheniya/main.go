package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		// i = 1
		fmt.Printf("Таблица умножения на %d\n", i)

		for j := 1; j <= 10; j++ {
			fmt.Printf("%d * %d = %d\n", i, j, i*j) // 1
		}
		// j = 1
		fmt.Println()
	}
}

// Как выйти сразу из двух вложенных циклов?
// outer:
// 	for i := 1; i <= 3; i++ {

// 		for j := 1; j <= 3; j++ {

// 			if i == 2 && j == 2 {
// 				break outer
// 			}

// 			fmt.Println(i, j)
// 		}
// 	}
// 	//break outer завершит сразу внешний цикл.

// }
