package main

import "fmt"

func main() {

	var word string

	fmt.Print("Введите слово: ")
	fmt.Scanln(&word)

	letters := make(map[rune]int)

	for _, ch := range word {
		letters[ch]++
	}

	for ch, count := range letters {
		fmt.Println(string(ch), "→", count)
	}
}
