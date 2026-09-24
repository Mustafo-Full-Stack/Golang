package main

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound     = errors.New("запись не найдена")
	ErrUnauthorized = errors.New("нет доступа")
	ErrInvalidInput = errors.New("некорректные данные")
)

var users = map[int]string{1: "Али", 2: "Вера"}

func getUser(id int) (string, error) {
	if id <= 0 {
		return "", ErrInvalidInput
	}
	name, ok := users[id]
	if !ok {
		return "", ErrNotFound
	}

	return name, nil
}

func main() {
	for _, id := range []int{1, 99, -1} {
		name, err := getUser(id)
		switch {
		case errors.Is(err, ErrNotFound):
			fmt.Printf("id=%d: пользователя нет, покажем пустую страницу\n", id)
		case errors.Is(err, ErrInvalidInput):
			fmt.Printf("id=%d: неверный запрос\n", id)
		case err != nil:
			fmt.Printf("id=%d: неизвестная ошибка: %v\n", id, err)
		default:
			fmt.Printf("id=%d: нашли %s\n", id, name)
		}
	}
}
