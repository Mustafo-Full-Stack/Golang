package wallet

import "errors"

type Wallet struct {
	balance float64
}

// New (функция-конструктор) создает новый экземпляр Wallet.
// Начальный баланс кошелька равен 0.
func New() *Wallet {
	return &Wallet{}
}

func (w *Wallet) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	w.balance += amount
	return nil
}

func (w *Wallet) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	if amount > w.balance {
		return errors.New("insufficient balance")
	}

	w.balance -= amount

	return nil
}

func (w *Wallet) Balance() float64 {
	return w.balance
}
