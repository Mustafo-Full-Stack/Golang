package main

import "fmt"

func main() {
	counts := make(map[string]int)

	// главная красота map в Go: не нужно проверять "а есть ли уже такой ключ"
	// перед инкрементом — если ключа нет, Go сам создаёт его со значением 0,
	// и сразу прибавляет 1
	counts["go"]++   // ключа не было -> создался с 0 -> стал 1
	counts["go"]++   // 2
	counts["rust"]++ // 1

	fmt.Println(counts) // map[go:2 rust:1]
}
