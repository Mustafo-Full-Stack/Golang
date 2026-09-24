/*
1. Опиши структуру `Wallet` с приватным полем `balance` (float64). Добавь методы `Deposit(amount float64) error` (ошибка, если сумма отрицательная), `Withdraw(amount float64) error` (ошибка, если не хватает денег) и `Balance() float64`. В main() сделай несколько операций подряд и выведи баланс после каждой, а на одной операции специально получи и выведи ошибку.
*/

package main

import (
	"fmt"
)

type Wallet struct {
	balance float64
}

func (w *Wallet) Deposit(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("Депозит отрицательную сумму нельзя")
	}
	w.balance += amount
	return nil
}

func (w *Wallet) Withdraw(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("ошибка, нельзя снять отрицательную сумму")
	}
	if amount > w.balance {
		return fmt.Errorf("ошибка, сумма не хватает")
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() float64 {
	return w.balance
}

func main() {
	wallet := &Wallet{}

	err := wallet.Deposit(-100)
	if err != nil {
		fmt.Println("Ошибка при депозите:", err)
	} else {
		fmt.Println("Баланс после депозита:", wallet.Balance())
	}

	err = wallet.Withdraw(50)
	if err != nil {
		fmt.Println("Ошибка при снятии:", err)
	} else {
		fmt.Println("Баланс после снятия:", wallet.Balance())
	}
}
