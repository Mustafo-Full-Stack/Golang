package main

import (
	"bufio"
	"fmt"
	"os"
)

func readLines(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		fmt.Printf("%3d | %s\n", lineNum, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка сканирования: %w", err)
	}
	fmt.Println("Всего строк:", lineNum)
	return nil
}

func main() {
	data := []byte("первая строка\nвторая строка\nтретья строка\nчетвёртая строка")
	os.WriteFile("lines.txt", data, 0644)

	if err := readLines("lines.txt"); err != nil {
		fmt.Println("Ошибка:", err)
	}
}
