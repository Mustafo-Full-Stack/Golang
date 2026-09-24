package account

type Depositor interface {
	Deposit(amount float64) error
}

type Withdrawer interface {
	Withdraw(amount float64) error
}

type BalanceProvider interface {
	Balance() float64
}

type Account interface {
	Depositor
	Withdrawer
	BalanceProvider
}
