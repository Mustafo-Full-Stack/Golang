package main

import "fmt"

// строку переводим в срез rune (не byte!), чтобы правильно
// переворачивать и кириллицу, и любые другие символы
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(reverse("Привет"))
	fmt.Println(reverse("Go"))
}
