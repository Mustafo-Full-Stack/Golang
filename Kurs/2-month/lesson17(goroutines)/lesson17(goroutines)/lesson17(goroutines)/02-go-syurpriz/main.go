package main

import (
	"fmt"
	"time"
)

func downloadFile(name string) {
	fmt.Printf("Начали скачивать %s\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("Скачали %s\n", name)
}

func main() {
	start := time.Now()

	// Добавили слово "go" - и получили сюрприз
	go downloadFile("file1.zip")
	go downloadFile("file2.zip")
	go downloadFile("file3.zip")

	// Вывода почти нет! Спроси группу: почему?
	// Ответ: main завершился раньше горутин - и программа умерла вместе с ними.
	time.Sleep(1 * time.Second)
	fmt.Println("Всего заняло:", time.Since(start)) // ~0 секунд, но ничего не скачалось
}
