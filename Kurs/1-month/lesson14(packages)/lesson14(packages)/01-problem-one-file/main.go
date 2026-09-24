package main

import (
	"fmt"
	"strings"
)

// Это демо-версия "монолита" из 30 строк. У юзера к концу
// проекта такой файл легко разрастается на 500-800 строк: валидация,
// расчёты и вывод свалены вместе, и непонятно, где что искать и как
// переиспользовать код в другом проекте.

func main() {
	name := "Али"
	email := "ali@mail.tj"
	price := 250.0
	quantity := 3

	// валидация
	if name == "" {
		fmt.Println("ошибка: имя не может быть пустым")
		return
	}
	if !strings.Contains(email, "@") {
		fmt.Println("ошибка: некорректный email")
		return
	}
	if quantity <= 0 {
		fmt.Println("ошибка: количество должно быть положительным")
		return
	}

	// расчёты
	total := price * float64(quantity)
	discount := 0.0
	if total > 500 {
		discount = total * 0.1
	}
	finalPrice := total - discount

	// вывод
	fmt.Printf("Покупатель: %s <%s>\n", name, email)
	fmt.Printf("Товаров: %d по %.2f\n", quantity, price)
	fmt.Printf("Сумма: %.2f, скидка: %.2f, итог: %.2f\n", total, discount, finalPrice)

	// как это тестировать отдельно от вывода?
	// Как использовать те же расчёты в другой программе, не копируя код?
	// Ответ - разложить по пакетам. Это и делаем в следующих шагах.
}
