package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field   string
	Value   any
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("поле %q (значение %v): %s", e.Field, e.Value, e.Message)
}

func validateUser(name string, age int) error {
	if name == "" {
		return &ValidationError{Field: "name", Value: name, Message: "не может быть пустым"}
	}
	if age < 18 {
		return &ValidationError{Field: "age", Value: age, Message: "должен быть не меньше 18"}
	}
	return nil
}

func main() {
	err := validateUser("", 25)
	fmt.Println(err)

	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Println("Проблемное поле:", ve.Field)
		fmt.Println("Значение:", ve.Value)
	}
}
