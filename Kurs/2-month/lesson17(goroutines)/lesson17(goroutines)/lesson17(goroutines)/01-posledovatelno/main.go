package main

import (
	"fmt"
	"time"
)

// Скачиваем 3 "файла" по очереди - каждый по 1 секунде.
// Вопрос группе после запуска: "Можно быстрее?"
func downloadFile(name string) {
	fmt.Printf("Начали скачивать %s\n", name)
	time.Sleep(1 * time.Second) // имитация сети
	fmt.Printf("Скачали %s\n", name)
}

func main() {
	start := time.Now()

	downloadFile("file1.zip")
	downloadFile("file2.zip")
	downloadFile("file3.zip")

	fmt.Println("Всего заняло:", time.Since(start)) // ~3 секунды
}
