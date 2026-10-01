package main

import (
	"fmt"
	"os"
)

func main() {
	os.WriteFile("test.txt", []byte("привет"), 0644)

	os.Rename("test.txt", "renamed.txt") // переименовать/переместить
	fmt.Println("Переименовали test.txt в renamed.txt")

	os.Mkdir("data", 0755)              // создать одну папку
	os.MkdirAll("data/2026/july", 0755) // создать всю цепочку папок
	fmt.Println("Создали папки data/2026/july")

	os.Remove("renamed.txt") // удалить файл
	fmt.Println("Удалили renamed.txt")

	// список файлов и папок в текущей директории
	entries, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("\nСодержимое папки:")
	for _, e := range entries {
		marker := "файл"
		if e.IsDir() {
			marker = "папка"
		}
		fmt.Printf("[%s] %s\n", marker, e.Name())
	}

	os.RemoveAll("data") // удалить папку рекурсивно вместе с содержимым
}
