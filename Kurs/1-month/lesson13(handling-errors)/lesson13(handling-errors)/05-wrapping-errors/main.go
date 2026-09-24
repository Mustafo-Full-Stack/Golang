package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("запись не найдена")

func findUser(id int) error {
	return ErrNotFound
}

func createOrder(userID int) error {
	if err := findUser(userID); err != nil {
		return fmt.Errorf("не удалось создать заказ для user %d: %w", userID, err)
	}
	return nil
}

func handleRequest(userID int) error {
	if err := createOrder(userID); err != nil {
		return fmt.Errorf("обработка запроса не удалась: %w", err)
	}
	return nil
}

func main() {
	err := handleRequest(99)
	fmt.Println("Полное сообщение:")
	fmt.Println(" ", err)

	if errors.Is(err, ErrNotFound) {
		fmt.Println("Первопричина: запись не найдена")
	}

	fmt.Println("\nЦепочка:")
	for e := err; e != nil; e = errors.Unwrap(e) {
		fmt.Println("  →", e)
	}
}
