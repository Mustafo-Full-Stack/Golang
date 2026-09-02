package main

import "fmt"

/*
Условие

Написать программу, которая получает:
имя студента;
количество баллов за экзамен.

Правила:
90 - 100  → Отлично (A)
75 - 89   → Хорошо (B)
60 - 74   → Удовлетворительно (C)
40 - 59   → Нужно подтянуть знания (D)
0 - 39    → Не сдал (F)

Дополнительно:
если балл больше 100 → ошибка;
если балл меньше 0 → ошибка.
*/

// первый вариант

func main() {
	// var score int

	// fmt.Print("Введите балл: ")
	// fmt.Scan(&score)

	// if score > 100 {
	// 	fmt.Println("Ошибка: максимум 100 баллов")
	// } else if score < 0 {
	// 	fmt.Println("Ошибка: балл не может быть отрицательным")
	// } else if score >= 90 {
	// 	fmt.Println("Отлично (A)")
	// } else if score >= 75 {
	// 	fmt.Println("Хорошо (B)")
	// } else if score >= 60 {
	// 	fmt.Println("Удовлетворительно (C)")
	// } else if score >= 40 {
	// 	fmt.Println("Нужно подтянуть знания (D)")
	// } else {
	// 	fmt.Println("Не сдал (F)")
	// }
	mainV1()
}

/* Но если у нас будет 15 оценок, 20 ролей, 30 статусов — что будет?

if ...
else if ...
else if ...
else if ...
else if ...
else if ...

тяжело читать((((((((((
*/

// best practice через switch case
func mainV1() {
	var score int

	fmt.Print("Введите балл: ")
	fmt.Scan(&score)

	if score < 0 || score > 100 {
		fmt.Println("Ошибка: неправильный балл")
		return
		//ниже не пойдет!
	}

	grade := score / 10

	switch grade {
	case 10, 9:
		fmt.Println("Отлично (A)")

	case 8, 7:
		fmt.Println("Хорошо (B)")

	case 6:
		fmt.Println("Удовлетворительно (C)")

	case 5, 4:
		fmt.Println("Нужно подтянуть знания (D)")

	default:
		fmt.Println("Не сдал (F)")
	}
}

// best practice через switch case
func calculateGrade(score int) string {
	switch score / 10 {
	case 10, 9:
		return "Отлично (A)"

	case 8, 7:
		return "Хорошо (B)"

	case 6:
		return "Удовлетворительно (C)"

	case 5, 4:
		return "Нужно подтянуть знания (D)"

	default:
		return "Не сдал (F)"
	}
}
