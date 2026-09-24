package main

import "fmt"

func main() {
	// в Go нет отдельного типа "множество" (set), его делают через map[Тип]bool
	seen := make(map[string]bool)

	visitors := []string{"Али", "Вера", "Али", "Тимур", "Вера"}
	for _, v := range visitors {
		seen[v] = true 
	}

	fmt.Println("Уникальных посетителей:", len(seen))
	for name := range seen {
		fmt.Print(name, " ")
	}
	fmt.Println()

	// это практичный приём для удаления дубликатов из среза
}
