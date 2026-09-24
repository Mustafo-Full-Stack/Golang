package bankaccount

import "errors"

type BankAccount struct {
	balance float64
}

// New(функция-конструктор) создает новый банковский счет.
// Начальный баланс счета равен 0.
func New() *BankAccount {
	return &BankAccount{}
}

func (b *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	b.balance += amount

	return nil
}

func (b *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	if amount > b.balance {
		return errors.New("insufficient balance")
	}

	b.balance -= amount

	return nil
}

func (b *BankAccount) Balance() float64 {
	return b.balance
}
