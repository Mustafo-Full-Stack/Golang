package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Правильно: пишем через bufio.Writer и обязательно делаем Flush.
	f, _ := os.Create("ok.txt")
	writer := bufio.NewWriter(f)
	fmt.Fprintln(writer, "данные")
	writer.Flush()
	f.Close()
	content, _ := os.ReadFile("ok.txt")
	fmt.Println("Файл записан правильно:", string(content))

	// поломка 1 - забыли Flush.
	// f, _ := os.Create("broken1.txt")
	// writer := bufio.NewWriter(f)
	// fmt.Fprintln(writer, "данные")
	// f.Close()
	// // файл будет пустым - данные остались в буфере и умерли вместе с программой

	// поломка 2 - забыли Close после Open.
	// f, _ := os.Open("ok.txt")
	// scanner := bufio.NewScanner(f)
	// _ = scanner
	// // в цикле на тысячи файлов приведёт к "too many open files" - кончатся
	// // дескрипторы, которые выделяет операционная система

	// поломка 3 - печать байтов без string().
	// content, _ := os.ReadFile("ok.txt")
	// fmt.Println(content) // [208 180 208 176 208 189 ...] вместо текста

	// поломка 4 - os.WriteFile вместо дозаписи стирает старое содержимое.
	// os.WriteFile("ok.txt", []byte("новая запись"), 0644)
	// // старое содержимое исчезло; для дозаписи нужен os.OpenFile с os.O_APPEND

	// поломка 5 - не проверили ошибку open, работаем с nil-файлом.
	// f, _ := os.Open("не_существует.txt")
	// defer f.Close()
	// scanner := bufio.NewScanner(f)
	// _ = scanner
	// // panic: runtime error: invalid memory address or nil pointer dereference
}
