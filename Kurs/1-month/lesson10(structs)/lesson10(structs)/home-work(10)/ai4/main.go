package main

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance float64
}

func (w *Wallet) Deposit(amout float64) error {
	fmt.Println("Текущий Баланс:", w.balance)

	if amout < 0 {
		return errors.New("сумма не может быть отрицательной")
	}

	w.balance += amout
	fmt.Println("Попольнение:", w.balance)

	return nil
}

func (w *Wallet) Withdraw(amout float64) error {

	if amout > w.balance {
		return errors.New("Ошибка ! Денег не достатечно")
	}

	err := w.Deposit(100)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("ВЫВОД: -", amout)
	w.balance -= amout
	fmt.Println("Баланс:", w.balance)

	return nil
}

func main() {

	user := Wallet{
		balance: 500,
	}

	user.Deposit(200)

	user.Withdraw(1000)
}
