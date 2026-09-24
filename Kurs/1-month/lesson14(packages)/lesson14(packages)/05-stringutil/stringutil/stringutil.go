package stringutil

import "strings"

// Reverse переворачивает строку (корректно для Unicode).
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome проверяет палиндром без учёта регистра и пробелов.
func IsPalindrome(s string) bool {
	clean := strings.ToLower(strings.ReplaceAll(s, " ", ""))
	return clean == Reverse(clean)
}

// Truncate обрезает строку до n символов, добавляя многоточие.
func Truncate(s string, n int) string {
	runes := []rune(s)
	if len(runes) <= n {
		return s
	}
	return string(runes[:n]) + "..."
}

// CountWords считает слова.
func CountWords(s string) int {
	return len(strings.Fields(s))
}
