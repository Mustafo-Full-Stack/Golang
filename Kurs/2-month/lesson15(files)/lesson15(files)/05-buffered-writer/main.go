package main

import (
	"bufio"
	"fmt"
	"os"
)

func writeManyLines(filename string, count int) error {
	f, err := os.Create(filename) // создаёт файл или очищает существующий
	if err != nil {
		return err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	defer writer.Flush() // ОБЯЗАТЕЛЬНО - иначе часть строк останется в буфере

	for i := 1; i <= count; i++ {
		fmt.Fprintf(writer, "Строка номер %d\n", i)
	}
	return nil
}

func main() {
	if err := writeManyLines("big.txt", 1000); err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Записано 1000 строк")

	content, _ := os.ReadFile("big.txt")
	fmt.Println("Размер файла:", len(content), "байт")
}
