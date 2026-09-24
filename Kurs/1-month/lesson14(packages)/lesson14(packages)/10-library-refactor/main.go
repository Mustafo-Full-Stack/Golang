package main

import (
	"fmt"

	"library/models"
	"library/storage"
)

// main.go - только сценарий использования: создать библиотеку,
// наполнить книгами, выдать/принять и вывести результат.
// Вся логика живёт в models и storage - именно это и даёт разбиение
// на пакеты: main.go остаётся коротким и понятным.
func main() {
	lib := storage.New()

	bookOne, err := models.NewBook("Мастер и Маргарита", "Булгаков", 1967)
	if err != nil {
		return
	}

	bookTwo, err := models.NewBook("Преступление и наказание", "Достоевский", 1866)
	if err != nil {
		return
	}

	// lib.Add(models.NewBook("Мастер и Маргарита", "Булгаков", 1967))
	// lib.Add(models.NewBook("Преступление и наказание", "Достоевский", 1866))
	lib.Add(bookOne)
	lib.Add(bookTwo)

	if err := lib.Borrow("Мастер и Маргарита"); err != nil {
		fmt.Println("ошибка:", err)
	}

	if err := lib.Borrow("Мастер и Маргарита"); err != nil {
		fmt.Println("ошибка:", err) // уже на руках
	}

	if err := lib.Borrow("Незнакомая книга"); err != nil {
		fmt.Println("ошибка:", err) // не найдена
	}

	if err := lib.Return("Мастер и Маргарита"); err != nil {
		fmt.Println("ошибка:", err)
	}

	fmt.Println("\nВся библиотека:")
	for _, b := range lib.All() {
		fmt.Println(" -", b)
	}
}
