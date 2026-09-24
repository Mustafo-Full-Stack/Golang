package main

import (
	"fmt"
	"time"
)

// Урок 11, Шаг 11: мини-проект - банковский счёт с методами

type Transaction struct {
	Type   string
	Amount float64
	Time   time.Time
}

type Account struct {
	Owner   string
	balance float64 // приватное поле - менять можно только через методы
	history []Transaction
}

func NewAccount(owner string, initial float64) *Account {
	a := &Account{Owner: owner, history: []Transaction{}}
	if initial > 0 {
		a.Deposit(initial)
	}
	return a
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("сумма должна быть положительной")
	}
	a.balance += amount
	a.history = append(a.history, Transaction{"Пополнение", amount, time.Now()})
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("сумма должна быть положительной")
	}
	if amount > a.balance {
		return fmt.Errorf("недостаточно средств: есть %.2f, нужно %.2f", a.balance, amount)
	}
	a.balance -= amount
	a.history = append(a.history, Transaction{"Снятие", amount, time.Now()})
	return nil
}

func (a Account) Balance() float64 {
	return a.balance
}

func (a Account) String() string {
	return fmt.Sprintf("Счёт %s: %.2f (операций: %d)", a.Owner, a.balance, len(a.history))
}

func main() {
	acc := NewAccount("Али", 1000)
	acc.Deposit(500)

	if err := acc.Withdraw(5000); err != nil {
		fmt.Println("Ошибка:", err) // Ошибка: недостаточно средств: есть 1500.00, нужно 5000.00
	}

	acc.Withdraw(300)
	fmt.Println(acc) // Счёт Али: 1200.00 (операций: 3)

	for _, t := range acc.history {
		fmt.Printf("  %s: %.2f\n", t.Type, t.Amount)
	}
}

// здесь сходится весь урок: конструктор, pointer-receiver, приватные поля, String(), ошибки
