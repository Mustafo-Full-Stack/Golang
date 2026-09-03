/*
Пользователь вводит число N.
Найдите сумму чисел от 1 до N.
Числа, которые делятся на 3,
необходимо пропускать.
Если встретилось число,
которое делится на 7,
цикл должен завершиться.
*/

package main

import (
	"fmt"
)

func main() {
	var N, sum int

	fmt.Println("Введите Число :")
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		if i%3 == 0 {
			continue
		}
		if i%7 == 0 {
			return
		}
		sum += i
	}
	fmt.Println("Сумма:", sum)
}
