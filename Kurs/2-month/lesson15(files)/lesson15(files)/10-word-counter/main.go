package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func analyzeFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("не удалось открыть: %w", err)
	}
	defer f.Close()

	lines, words, chars := 0, 0, 0
	freq := make(map[string]int)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lines++
		chars += len([]rune(line))
		for _, w := range strings.Fields(strings.ToLower(line)) {
			words++
			freq[strings.Trim(w, ".,!?;:")]++
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Printf("Строк:    %d\nСлов:     %d\nСимволов: %d\n", lines, words, chars)

	type kv struct {
		Word  string
		Count int
	}
	var pairs []kv
	for w, c := range freq {
		pairs = append(pairs, kv{w, c})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].Count > pairs[j].Count })

	fmt.Println("Топ-3 слова:")
	for i := 0; i < 3 && i < len(pairs); i++ {
		fmt.Printf("  %s — %d\n", pairs[i].Word, pairs[i].Count)
	}
	return nil
}

func main() {
	text := "Go это язык программирования.\n" +
		"Go простой и быстрый язык.\n" +
		"Программирование на Go приносит удовольствие."
	os.WriteFile("article.txt", []byte(text), 0644)

	if err := analyzeFile("article.txt"); err != nil {
		fmt.Println("Ошибка:", err)
	}
}
