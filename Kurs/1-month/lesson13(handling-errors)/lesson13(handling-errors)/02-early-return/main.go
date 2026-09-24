package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID   int
	Name string
}

type Order struct {
	ID     int
	UserID int
}

var users = map[int]User{
	1: {ID: 1, Name: "Алишер"},
}

func findUser(id int) (User, error) {
	user, ok := users[id]
	if !ok {
		return User{}, errors.New("пользователь не найден")
	}
	return user, nil
}

func createOrder(user User) (Order, error) {
	return Order{ID: 100, UserID: user.ID}, nil
}

func sendEmail(order Order) error {
	fmt.Printf("письмо о заказе #%d отправлено\n", order.ID)
	return nil
}

func processOrder(id int) error {
	user, err := findUser(id)
	if err != nil {
		return err // сразу выходим
	}

	order, err := createOrder(user)
	if err != nil {
		return err
	}

	if err := sendEmail(order); err != nil {
		return err
	}

	fmt.Println("Заказ обработан")
	return nil
}

func main() {
	if err := processOrder(1); err != nil {
		fmt.Println("Ошибка:", err)
	}

	if err := processOrder(99); err != nil {
		fmt.Println("Ошибка:", err)
	}
}
