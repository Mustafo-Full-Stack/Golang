package main

import (
	"fmt"
	"strings"
)

// именованные результаты - переменные результата объявлены прямо
// в скобках после параметров, уже с нулевыми значениями,
// и их можно сразу использовать внутри функции

func rectangle(w, h float64) (area, perimeter float64) {
	area = w * h
	perimeter = 2 * (w + h)
	return // "голый" return - вернёт то, что уже лежит в area и perimeter
}

func splitName(full string) (first, last, given string) {
	parts := strings.Fields(full)
	if len(parts) > 0 {
		first = parts[0]
	}
	if len(parts) > 1 {
		last = parts[1]
	}

	if len(parts) > 2 {
		given = parts[2]
	}

	return
}

func main() {
	area, perimeter := rectangle(3, 4)
	fmt.Printf("Площадь %.1f, периметр %.1f\n", area, perimeter)
	fmt.Println(splitName("АлиРахимов Рахимович"))

	// плюс: имена в сигнатуре сразу документируют, что возвращает функция
	// минус: "голый" return в длинной функции читать сложнее -
	// используй именованные результаты в коротких функциях
}
