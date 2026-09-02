package main

import "fmt"

/* Форма 2 — как while (только условие):
for условие {
	// тело
}

В других языках это `while`. В Go отдельного `while` нет.
*/

func main() {
	// balance := 1000
	// month := 0
	// for balance > 0 {
	// 	balance -= 300
	// 	//balance = balance - 300
	// 	month++
	// 	fmt.Printf("Месяц %d: осталось %d\n", month, balance)
	// }
	// fmt.Println("Деньги кончились через", month, "мес.")

	//Первые 15 чисел, кратных 7
	count := 0
	number := 1
	for count <= 15 {
		if number%7 == 0 {
			fmt.Printf("%d ", number)
			count++
		}
		number++
	}
	fmt.Println()

}
