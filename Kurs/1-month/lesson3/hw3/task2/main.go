package main

import (
	"fmt"
)

func main() {
	var salary float64

	fmt.Println("Введите размер заработной платы: ")
	fmt.Scan(&salary)

	if salary <= 0 {
		fmt.Println("Ошибка: зарплата должна бить больше 0")
		return
	}

	var Nalog float64

	switch {
	case salary <= 3000:
		Nalog = 0.05
	case salary <= 10000:
		Nalog = 0.10
	default:
		Nalog = 0.15
	}

	ResultNalog := salary * Nalog
	ResultSalary := salary - ResultNalog
	ResultNalog = Nalog * 100

	

	fmt.Printf("Сумма налога: %.2f сомони\n", ResultNalog)
	fmt.Printf("Сумма после налога: %.2f сомони\n", ResultSalary)
}
