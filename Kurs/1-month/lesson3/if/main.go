package main

import "fmt"

func main() {

	// // =====================================
	// // 1. Простое условие if else
	// // Совершеннолетие
	// // =====================================

	// // var age int

	// // fmt.Print("Введите возраст: ")
	// // fmt.Scan(&age)

	// // if age >= 18 { // 19 >= 18 - true
	// // 	fmt.Println("Совершеннолетний")
	// // } else {
	// // 	fmt.Println("Несовершеннолетний")
	// // }

	// // =====================================
	// // 2. Несколько условий if else if
	// // Оценка студента
	// // =====================================

	// var score int

	// fmt.Print("Введите балл: ")
	// fmt.Scan(&score)

	// if score >= 90 {
	// 	fmt.Println("Отлично")
	// } else if score >= 75 {
	// 	fmt.Println("Хорошо")
	// } else if score >= 60 {
	// 	fmt.Println("Удовлетворительно")
	// } else {
	// 	fmt.Println("Неудовлетворительно")
	// }

	// =====================================
	// 3. Несколько независимых if
	// Проверка условий полета
	// =====================================

	// =====================================
	// 3. Несколько независимых if
	// Проверка условий полета
	// =====================================

	// var (
	// 	passengerAge int // 0
	// 	hasPassport  bool // false
	// 	hasTicket    bool // false
	// )

	// fmt.Print("Введите возраст пассажира: ")
	// fmt.Scan(&passengerAge)

	// // fmt.Print("Есть паспорт? (true/false): ")
	// // fmt.Scan(&hasPassport)
	
	// fmt.Print("Есть билет? (true/false): ")
	// fmt.Scan(&hasTicket)

	// if passengerAge >= 18 {
	// 	fmt.Println("Возраст подходит, и паспорт имеется))")
	// 	hasPassport = true
	// } else {
	// 	fmt.Println("Возраст не подходит, и паспорта нет значит!))")
	// 	hasPassport = false
	// }

	// // if !hasPassport { //hasPassport = false !true = false, !false=true
	// // 	fmt.Println("Нет паспорта")
	// // }
	
	// if hasPassport {
	// 	fmt.Println(" Ест паспорта")
	// }

	// if !hasTicket {
	// 	fmt.Println("Нужно купить билет")
	// }

	// =====================================
	// 4. Логические операторы
	// И (&&), ИЛИ (||)
	// =====================================
	var passengerAge int

	if passengerAge < 7 || passengerAge > 65 {
		fmt.Println("Льготный тариф")
	} else {
		fmt.Println("Полный тариф")
	}

	// =====================================
	// 5. Проверка диапазона
	// =====================================

	temp := 25.1

	if temp >= 18 && temp <= 25 {
		fmt.Println("Комфортная температура")
	} else {
		fmt.Println("Температура некомфортная")
	}
}
