package main

import (
	"fmt"

	"strutil-homework/strutil"
)

// Домашнее задание 2:
// Проверь несколько слов и фраз. Для каждой строки выведи её разворот
// и укажи, является ли она палиндромом.

func main() {
	var (
		word1 = "level"
		word2 = "топор"
		word3 = "привет"
	)
	fmt.Println(
		strutil.Reverse(word1),
		strutil.IsPalindrome(word1),
	)
	fmt.Println(
		strutil.Reverse(word2),
		strutil.IsPalindrome(word2),
	)
	fmt.Println(
		strutil.Reverse(word3),
		strutil.IsPalindrome(word3),
	)
}

/*
Добавь ещё пару значений,
например "топот" и "привет",
и для каждого выведи исходную строку,
результат Reverse и ответ IsPalindrome.
Тогда увидишь, что переворот работает и с
кириллицей.
*/
