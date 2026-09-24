package main

import (
	"fmt"
)

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b BankAccount) Deposit(amout float64) {
	var text string
	var sum float64

	fmt.Println("Баланс:", b.Balance)
	fmt.Println(
		"Выберите опций:",
	)

	fmt.Println("<<< Вывод")
	fmt.Println("<<< Попольнение")

	fmt.Println("ПОЖАЛЮСТА НАПИШИТЕ вывод или попольнение с маленьким буквами для выбора опций!")
	fmt.Scanln(&text)

	switch text {
	case "вывод":
		fmt.Println("Введите Сумму для вывода:")
		fmt.Scanln(&sum)

		if sum > amout {
			fmt.Println("ОШИБКА ! Сумма для вывода Больше Баланса !")
		} else if sum <= amout {
			b.Balance -= amout

			fmt.Println("Баланс Успешно Выведен !")
			fmt.Println("Баланс:", b.Balance)
		}

	case "попольнение":
		fmt.Println("Введите Сумму для попольнение:")
		fmt.Scanln()

	default:
		fmt.Println("ОШИБКА ! ПОЖАЛЮСТА НАПИШИТЕ вывод или попольнение с маленьким буквами для выбора опций!")
		return
	}

	// if text == "вывод" {
	// 	fmt.Println("Введите Сумму для вывода:")
	// } else if text == "попольнение" {
	// 	fmt.Println("Введите Сумму для попольнение:")
	// } else {
	// 	fmt.Println("ОШИБКА ! ПОЖАЛЮСТА НАПИШИТЕ вывод или попольнение с маленьким буквами для выбора опций!")
	// }

}

func main() {
	user := BankAccount{
		Balance: 20,
	}

	user.Deposit(30)
}

//Давомша бад мекнм
