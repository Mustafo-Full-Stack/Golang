package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"Али":   25,
		"Вера":  30,
		"Тимур": 19,
	}

	// стандартный рецепт Go для стабильного вывода map:
	// 1. собираем все ключи в срез
	keys := make([]string, 0, len(ages))
	for k := range ages {
		keys = append(keys, k)
	}

	// 2. сортируем срез ключей
	sort.Strings(keys)

	// 3. идём по отсортированным ключам и берём значения из map
	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, ages[k])
	}
}
