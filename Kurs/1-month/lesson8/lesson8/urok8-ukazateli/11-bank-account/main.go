package main

import "fmt"

// Урок 8, Шаг 11: мини-проект - банковский счёт,
// реалистичный пример, где указатель действительно нужен

func deposit(balance *float64, amount float64) bool {
	if amount <= 0 {
		return false
	}
	*balance += amount
	return true
}

func withdraw(balance *float64, amount float64) bool {
	if amount <= 0 || amount > *balance {
		return false
	}
	*balance -= amount
	return true
}

func main() {
	balance := 1000.0
	amount := 222.2

	updatedBalance := depositWithPtr(&balance, &amount)

	fmt.Println(*updatedBalance)

	//balance = depositNew(balance, 222.22)

	//fmt.Println("После пополнения:balance", balance) // 1000.0

	//fmt.Println("После пополнения:newBalance", newBalance) // 1222.2

	// if !withdraw(&balance, 5000) {
	// 	fmt.Println("Недостаточно средств")
	// }

	// if withdraw(&balance, 300) {
	// 	fmt.Printf("Снятие. Баланс: %.2f\n", balance)
	// }
}

func depositNew(balance, amount float64) float64 {
	if amount <= 0 {
		return 0
	}
	balance += amount
	return balance
}

func depositWithPtr(balance, amount *float64) *float64 {
	if *amount <= 0 {
		return nil
	}
	*balance = *balance + *amount
	return balance
}
