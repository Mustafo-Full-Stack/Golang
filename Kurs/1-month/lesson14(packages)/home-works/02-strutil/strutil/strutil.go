package strutil

// Задание 2:
// Реализуй Reverse(s string) string для разворота строки.
// Реализуй IsPalindrome(s string) bool, используя Reverse.
// Учитывай, что строка может содержать Unicode-символы.
func Reverse(s string) string {
	ReturnedString := ""
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		ReturnedString += string(runes[i])
	}
	return ReturnedString
}

func IsPalindrome(s string) bool {
	if Reverse(s) == s {
		return true
	}
	return false
}
