package main

import (
	"fmt"
	"sort"
	"strings"
)

// Проект: подсчёт частоты слов в тексте

func main() {
	text := "go это просто go это быстро go это надёжно"

	words := strings.Fields(text) // разбиваем строку по пробелам на срез слов
	counts := make(map[string]int)

	for _, w := range words {
		counts[strings.ToLower(w)]++ // приводим к нижнему регистру и считаем
	}

	// сортируем ключи, чтобы вывод был одинаковым при каждом запуске
	keys := make([]string, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Println("Частота слов:")
	for _, k := range keys {
		fmt.Printf("%-10s %d\n", k, counts[k]) // %-10s - выравнивание по левому краю, поле 10 символов
	}

	// находим самое частое слово - просто проходим по map и ищем максимум
	maxWord, maxCount := "", 0
	for w, c := range counts {
		if c > maxCount {
			maxWord, maxCount = w, c
		}
	}
	fmt.Printf("Чаще всего: %q (%d раз)\n", maxWord, maxCount)
}
