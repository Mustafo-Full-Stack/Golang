package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Printf("Ошибка:%v", err)
	} else {
		fmt.Printf("Результат:%v", result)
	}

	// result, err = divide(10, 0)
	// if err != nil {
	// 	fmt.Println("Ошибка:", err)
	// } else {
	// 	fmt.Println("Результат:", result)
	// }
}
