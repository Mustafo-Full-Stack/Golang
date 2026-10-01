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

	go downloadFile("file1.zip")
	go downloadFile("file2.zip")
	go downloadFile("file3.zip")

	time.Sleep(2 * time.Second) // ждём вручную - это костыль!
	fmt.Println("Всего заняло:", time.Since(start)) // ~2 секунды вместо 3

	// Проблема: откуда мы знаем, сколько ждать?
	// Если файл качается 5 секунд - мы его потеряем.
	// Нужен нормальный способ ждать горутины -> sync.WaitGroup
}
