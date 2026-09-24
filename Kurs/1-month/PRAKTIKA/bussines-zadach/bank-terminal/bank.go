package main

import (
	"errors"
	"fmt"
)

type Account struct {
	name    string
	balance float64
	// active  bool
}

func (a *Account) Deposit(amout float64) error {
	if amout <= 0 {
		return errors.New("⚠️ Ошибка ! Сумма для попольнение должен быть больше 0")
	}

	a.balance += amout

	return nil
}

func (w *Account) Withdraw(amout float64) error {

	if amout <= 0 {
		return errors.New("⚠️ Ошибка ! Сумма для вывода должен быть больше 0")
	}

	w.balance -= amout

	return nil
}

func (b *Account) Balance() float64 {
	return b.balance
}

func main() {

	var NameUser string
	var BalanceUser float64
	var Options int
	var Vibor float64

	fmt.Println("==========")
	fmt.Println("ДОБРО ПОЖАЛОВАТЬ !")
	fmt.Println("==========")

	fmt.Println("")

	fmt.Println("Создайте Акаунт для продолжение!")
	fmt.Println("Введите Имя:")
	fmt.Scanln(&NameUser)

	fmt.Println("Введите Баланс:")
	fmt.Println("Баланс не можеть быть меньше 0 !")
	fmt.Scanln(&BalanceUser)

	if BalanceUser < 0 {
		fmt.Println("⚠️: БАЛАНС НЕ МОЖЕТ БЫТЬ МЕНЬШЕ 0 !")
		return
	}

	fmt.Println("==========")
	fmt.Println("РЕГИСТРАЦИЯ ПРОШЛО УСПЕШНО ✅ !")
	fmt.Println("==========")

	fmt.Println("")

	fmt.Println("==========")
	fmt.Println("ДОБРО ПОЖАЛОВАТЬ", NameUser, "В МИНИ-БАНК 🏦!")
	fmt.Println("==========")

	fmt.Println("")

	fmt.Println("Ваш Баланс:", BalanceUser)

	user := &Account{
		name:    NameUser,
		balance: BalanceUser,
	}

	for {
		fmt.Println("Введите операцию:")

		fmt.Println("1 - Пополнить")
		fmt.Println("2 - Вывод")
		fmt.Println("3 - Проверить баланс")
		fmt.Println("4 - Выход")
		fmt.Scanln(&Options)

		if Options == 1 {
			fmt.Println("Введите Сумму для Пополнения:")
			fmt.Scanln(&Vibor)

			fmt.Println("")

			if err := user.Deposit(Vibor); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("УСПЕШНО ПОПОЛНЕНО ✅")
			}

			fmt.Println("")
			fmt.Println("Ваш Баланс:", user.balance)
		}

		if Options == 2 {
			fmt.Println("Введите Сумму для вывода:")
			fmt.Scanln(&Vibor)

			fmt.Println("")

			if err := user.Withdraw(Vibor); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("УСПЕШНО ВЫВЕДЕНО ✅")
			}

			fmt.Println("")
			fmt.Println("Ваш Баланс:", user.balance)
		}

		if Options == 3 {
			fmt.Println("Обновляем Баланс. . .")

			fmt.Println("")
			fmt.Println("Ваш Баланс:", user.balance)
		}

		if Options == 4 {
			fmt.Println("Выход из программа . . ")
			fmt.Println("✅")
			return
		}
	}
}
