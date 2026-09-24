package main

import "fmt"

// Урок 8, Шаг 8: указатель как способ сказать "значения нет" -
// честнее, чем возвращать 0

func findAge(name string) *int {
	ages := map[string]int{"Али": 25, "Вера": 30}
	if age, ok := ages[name]; ok {
		return &age
	}
	return nil // "не найдено" - честно, а не 0
}

func main() {
	if age := findAge("Али"); age != nil {
		fmt.Println("Возраст:", *age)
	} else {
		fmt.Println("Не найден")
	}

	if age := findAge("Джон"); age == nil {
		fmt.Println("Джон не найден")
	}
}

// свяжи с уроком 6: это второй способ отличить "нет значения" от "значение 0"
