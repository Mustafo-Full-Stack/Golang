package main

import (
	"fmt"
)

func main() {
	var (
		UserBalance     int
		BalanceWithdraw int
	)

	fmt.Println("Введите Ваш Баланс : ")
	fmt.Scan(&UserBalance)

	fmt.Println("Введите сумму для вывода : ")
	fmt.Scan(&BalanceWithdraw)

	if BalanceWithdraw <= 0 {
		fmt.Println("Ошибке при снятие, попробуйте ещё раз !")
	} else if BalanceWithdraw > UserBalance {
		fmt.Println("Недостаточно средств")
	} else {
		fmt.Println("Успешно выведено ✅ !")
		fmt.Println("Ваш баланс : ", UserBalance-BalanceWithdraw)
	}
	if UserBalance <= 0 {
		fmt.Println("Баланс должен быть больше : 0")
	}
}
