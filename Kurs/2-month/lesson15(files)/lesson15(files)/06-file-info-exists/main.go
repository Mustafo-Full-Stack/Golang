package main

import (
	"fmt"
	"os"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func fileInfo(path string) {
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Имя:       ", info.Name())
	fmt.Println("Размер:    ", info.Size(), "байт")
	fmt.Println("Изменён:   ", info.ModTime().Format("02.01.2006 15:04"))
	fmt.Println("Директория:", info.IsDir())
	fmt.Println("Права:     ", info.Mode())
}

func main() {
	os.WriteFile("test.txt", []byte("привет"), 0644)

	fmt.Println("test.txt существует:", fileExists("test.txt"))
	fmt.Println("nope.txt существует:", fileExists("nope.txt"))

	fmt.Println()
	fileInfo("test.txt")
}
