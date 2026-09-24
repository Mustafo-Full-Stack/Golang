package main

import "fmt"

// safeDivide превращает панику от деления int на ноль в обычную ошибку.
// Границу вроде этой обычно ставят там, где ненадёжный код нельзя переписать,
// но и обрушивать всю программу из-за него нельзя (например, обработчик запроса).
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("восстановились после паники: %v", r)
		}
	}()

	result = a / b
	return
}

// unstableService имитирует нестабильный сервис: первые попытки проваливаются,
// а с третьей - отвечает успешно. Используем замыкание, чтобы функция сама
// помнила номер попытки между вызовами.
func unstableService() func() error {
	attempt := 0
	return func() error {
		attempt++
		if attempt < 3 {
			return fmt.Errorf("попытка %d: сервис недоступен", attempt)
		}
		return nil
	}
}

func retry(attempts int, f func() error) error {
	var err error
	for i := 1; i <= attempts; i++ {
		err = f()
		if err == nil {
			return nil
		}
		fmt.Println("  ", err)
	}
	return fmt.Errorf("не удалось выполнить операцию за %d попыток: %w", attempts, err)
}

func main() {
	fmt.Println("--- recover на границе: делим на ноль без падения программы ---")
	pairs := [][2]int{{10, 2}, {8, 0}, {20, 4}}
	for _, p := range pairs {
		result, err := safeDivide(p[0], p[1])
		if err != nil {
			fmt.Printf("%d / %d → ошибка: %v\n", p[0], p[1], err)
			continue
		}
		fmt.Printf("%d / %d = %d\n", p[0], p[1], result)
	}

	fmt.Println("\n--- retry: повторяем нестабильную операцию ---")
	call := unstableService()
	if err := retry(5, call); err != nil {
		fmt.Println("итог:", err)
	} else {
		fmt.Println("итог: операция выполнена успешно")
	}

	fmt.Println("\n--- ещё раз retry, но попыток не хватает ---")
	tooFewAttempts := unstableService()
	if err := retry(2, tooFewAttempts); err != nil {
		// итоговая ошибка обёрнута через %w, поэтому первопричина
		// ("сервис недоступен") видна и внутри неё, а не только в тексте.
		fmt.Println("итог:", err)
	}
}
