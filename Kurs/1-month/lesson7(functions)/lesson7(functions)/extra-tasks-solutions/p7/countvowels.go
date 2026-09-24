package main

import (
	"fmt"
	"strings"
)

// гласные буквы русского и английского алфавита, строчные и заглавные
const vowels = "aeiouAEIOUаеёиоуыэюяАЕЁИОУЫЭЮЯ"

func countVowels(s string) int {
	count := 0
	for _, letter := range s { // range по строке отдаёт руны, а не байты
		if strings.ContainsRune(vowels, letter) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countVowels("Привет")) // 2
	fmt.Println(countVowels("Hello"))  // 2
}
