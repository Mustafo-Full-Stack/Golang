package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("запись не найдена")

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль")
	}
	return a / b, nil
}

func main() {
	// Правильно: проверяем ошибку перед использованием результата.
	if result, err := divide(10, 0); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}

	// поломка 1 - игнорирование ошибки через _.
	// result, _ := divide(10, 0)
	// fmt.Println(result * 2) // 0, и никто не узнал о проблеме
	// "_" для ошибки - почти всегда красный флаг на код-ревью.

	// поломка 2 - проверка == nil вместо != nil.
	// if err == nil {
	//     fmt.Println("Ошибка:", err) // никогда не выполнится корректно
	// }

	// поломка 3 - сравнение обёрнутой ошибки через == вместо errors.Is.
	wrapped := fmt.Errorf("контекст: %w", ErrNotFound)
	fmt.Println(wrapped == ErrNotFound)          // false!
	fmt.Println(errors.Is(wrapped, ErrNotFound)) // true

	// поломка 4 - %v вместо %w рвёт цепочку обёртывания.
	lost := fmt.Errorf("контекст: %v", ErrNotFound)
	fmt.Println(errors.Is(lost, ErrNotFound)) // false - цепочка потеряна

	// поломка 5 - паника вместо ошибки для ожидаемой ситуации.
	// nums := []int{1, 2, 3}
	// fmt.Println(nums[10])
	// panic: runtime error: index out of range [10] with length 3
	//
	// Выход за границы среза - ошибка программиста, паника здесь уместна.
	// А вот "файл не найден" или "неверный ввод" - ожидаемые ситуации,
	// для них нужен error, а не panic.
}
