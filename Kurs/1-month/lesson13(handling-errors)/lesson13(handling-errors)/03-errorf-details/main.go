package main

import "fmt"

// validateAge - validate input age
func validateAge(age int) error {
	// if age < 0 {
	// 	return fmt.Errorf("возраст не может быть отрицательным: %d", age)
	// }
	// if age > 150 {
	// 	return fmt.Errorf("возраст слишком большой: %d", age)
	// }
	if age < 0 || age > 150 {
		return fmt.Errorf("не валидный возраст: %v", age)
	}
	return nil
}

func main() {
	for _, age := range []int{25, -5, 200} {
		if err := validateAge(age); err != nil {
			fmt.Println("Ошибка валидации:", err)
		} else {
			fmt.Println("Возраст", age, "— корректен")
		}
	}
}
