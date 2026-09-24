package main

import "fmt"

func main() {
	// значением map может быть что угодно, в том числе срез
	// категория -> список товаров
	catalog := map[string][]string{
		"Фрукты": {"яблоко", "банан"},
		"Овощи":  {"морковь", "лук"},
	}

	// добавляем элемент в срез, который лежит внутри map
	catalog["Фрукты"] = append(catalog["Фрукты"], "вишня")

	for cat, items := range catalog {
		fmt.Printf("%s (%d): %v\n", cat, len(items), items)
	}

	// значением также может быть другая map или структура (структуры - на следующих уроках)
}
