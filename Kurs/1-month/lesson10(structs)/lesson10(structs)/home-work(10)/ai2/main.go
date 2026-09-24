package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b BankAccount) ShowBalance() {
	fmt.Println("Данные:")
	fmt.Println("Владелец:", b.Owner)
	fmt.Println("Баланс:", b.Balance, "$")
}

func main() {
	user := BankAccount{
		Owner:   "Mustafo",
		Balance: 500,
	}

	user.ShowBalance()
}
