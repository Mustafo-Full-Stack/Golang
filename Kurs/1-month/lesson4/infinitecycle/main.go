package main

import "fmt"

/*
Форма 3 — бесконечный:
for {
	// тело, выход через break
}

break и continue
break -Полностью выйти из цикла
continue - Пропустить остаток тела, перейти к следующей итерации
*/

func main() {
	// for {
	// 	fmt.Println("\n1 — Привет")
	// 	fmt.Println("2 — Пока")
	// 	fmt.Println("0 — Выход")
	// 	fmt.Print("Ваш выбор: ")

	// 	var choice int
	// 	fmt.Scanln(&choice)

	// 	if choice == 0 {
	// 		fmt.Println("До свидания!")
	// 		break
	// 	}

	// 	switch choice {
	// 	case 1:
	// 		fmt.Println("Привет!")
	// 	case 2:
	// 		fmt.Println("Пока!")
	// 	default:
	// 		fmt.Println("Неизвестный пункт")
	// 	}
	// }

	mainV1()
}

// continue
func mainV1() {
	// fmt.Println("Нечётные числа до 20:")
	// for i := 1; i <= 20; i++ {
	// 	if i%2 == 0 {
	// 		continue // чётные пропускаем
	// 	}
	// 	fmt.Print(i, " ")
	// }
	// fmt.Println()

	fmt.Println("Числа от 1 до 30, кратные 3, но не кратные 5:")
	for i := 1; i <= 30; i++ {
		if i%3 != 0 {
			continue
		}
		if i%5 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println()
}
