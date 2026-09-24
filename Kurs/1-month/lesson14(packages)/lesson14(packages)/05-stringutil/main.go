package main

import (
	"fmt"

	"myapp/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("привет"))
	fmt.Println(stringutil.IsPalindrome("А роза упала на лапу Азора"))
	fmt.Println(stringutil.Truncate("Изучаем пакеты в Go", 10))
	fmt.Println(stringutil.CountWords("Изучаем пакеты в Go"))

	// Почему Reverse работает через []rune, а не []byte:
	// в UTF-8 русская буква занимает 2 байта, а не 1.
	word := "привет"
	fmt.Println("len в байтах:", len(word))        // 12 - байты, а не буквы
	fmt.Println("len в рунах:", len([]rune(word))) // 6 - настоящее число букв

	broken := []byte(word)
	for i, j := 0, len(broken)-1; i < j; i, j = i+1, j-1 {
		broken[i], broken[j] = broken[j], broken[i]
	}
	fmt.Println("развернули по байтам (сломано):", string(broken))
}
