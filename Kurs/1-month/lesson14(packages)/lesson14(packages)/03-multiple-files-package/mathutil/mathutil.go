package mathutil

import "fmt"

// Add складывает два числа.
func Add(a, b int) int {
	return a + b
}

// Multiply умножает два числа.
func Multiply(a, b int) int {
	return a * b
}

// Max возвращает большее из двух чисел.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// validate - приватная, доступна из любого файла этого же пакета.
func validate(n int) bool {
	return n >= 0
}

// Factorial считает факториал.
func Factorial(n int) (int, error) {
	if !validate(n) {
		return 0, fmt.Errorf("факториал не определён для %d", n)
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}
