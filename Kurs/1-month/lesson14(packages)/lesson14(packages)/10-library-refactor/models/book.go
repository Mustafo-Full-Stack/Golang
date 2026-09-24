package models

import (
	"errors"
	"fmt"
)

// Book - модель книги. Пакет models ничего не знает о хранилище -
// только описывает данные и поведение самой книги.
type Book struct {
	Title    string
	Author   string
	Year     int
	Borrowed bool
}

// NewBook - конструктор.
func NewBook(title, author string, year int) (*Book, error) {
	if title == "" || author == "" {
		return nil, errors.New("пустой title или author ")
	}
	if year <= 0 {
		return nil, errors.New("не валидный год издания")
	}
	return &Book{Title: title, Author: author, Year: year}, nil
}

func (b *Book) String() string {
	status := "в наличии"
	if b.Borrowed {
		status = "на руках"
	}
	return fmt.Sprintf("%q (%s, %d) - %s", b.Title, b.Author, b.Year, status)
}
