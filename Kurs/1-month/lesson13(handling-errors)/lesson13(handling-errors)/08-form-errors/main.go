package main

import (
	"errors"
	"fmt"
	"strings"
)

type FormErrors []error

func (fe FormErrors) Error() string {
	msgs := make([]string, len(fe))
	for i, e := range fe {
		msgs[i] = e.Error()
	}
	return strings.Join(msgs, "; ")
}

func validateForm(name, email string, age int) error {
	var errs FormErrors

	if name == "" {
		errs = append(errs, errors.New("имя обязательно"))
	}
	if !strings.Contains(email, "@") {
		errs = append(errs, fmt.Errorf("некорректный email: %q", email))
	}
	if age < 18 {
		errs = append(errs, fmt.Errorf("возраст меньше 18: %d", age))
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}

func main() {
	if err := validateForm("", "bad-email", 15); err != nil {
		fmt.Println("Форма не прошла проверку:")
		if fe, ok := err.(FormErrors); ok {
			for i, e := range fe {
				fmt.Printf("  %d. %v\n", i+1, e)
			}
		}
	}
}
