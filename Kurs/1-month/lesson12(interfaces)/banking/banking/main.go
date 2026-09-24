package main

import (
	"banking/account"
	"banking/bankaccount"
	"banking/wallet"
	"fmt"
)

func main() {
	w := wallet.New()
	b := bankaccount.New()

	var (
		walletAccount account.Account = w
		bankAccount   account.Account = b
	)

	fmt.Println("--> Wallet <--")

	fmt.Printf("Initial balance: %.2f\n", walletAccount.Balance())

	fmt.Println("Deposit 500")

	if err := walletAccount.Deposit(500); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Balance after deposit: %.2f\n", walletAccount.Balance())

	fmt.Println("Withdraw 120")

	if err := walletAccount.Withdraw(120); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Balance after withdraw: %.2f\n\n", walletAccount.Balance())

	fmt.Println("--> Bank Account <--")

	fmt.Printf("Initial balance: %.2f\n", bankAccount.Balance())

	fmt.Println("Deposit 5000")

	if err := bankAccount.Deposit(5000); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Balance after deposit: %.2f\n", bankAccount.Balance())

	fmt.Println("Withdraw 1500")

	if err := bankAccount.Withdraw(1500); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Balance after withdraw: %.2f\n\n", bankAccount.Balance())

	fmt.Println("--> All Accounts <--")

	accounts := []account.Account{
		walletAccount,
		bankAccount,
	}

	for i, acc := range accounts {
		fmt.Printf("Account #%d before deposit: %.2f\n", i+1, acc.Balance())

		fmt.Println("Deposit 100")

		acc.Deposit(100)

		fmt.Printf("Account #%d after deposit: %.2f\n\n", i+1, acc.Balance())
	}

	fmt.Println("--> Error Examples <--")

	fmt.Printf("Current wallet balance: %.2f\n", walletAccount.Balance())

	fmt.Println("Trying deposit -100")

	if err := walletAccount.Deposit(-100); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("Balance after failed deposit: %.2f\n", walletAccount.Balance())

	fmt.Println("Trying withdraw 100000")

	if err := walletAccount.Withdraw(100000); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("Balance after failed withdraw: %.2f\n", walletAccount.Balance())
}
