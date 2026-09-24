package storage

import (
	"fmt"

	"library/models"
)

// Library хранит книги и умеет их искать, выдавать и принимать обратно.
// Пакет storage импортирует models (хранилище знает про модель),
// а не наоборот - это правильное направление зависимости.
type Library struct {
	books []*models.Book
}

// New - конструктор пустой библиотеки.
func New() *Library {
	return &Library{}
}

func (l *Library) Add(b *models.Book) {
	l.books = append(l.books, b)
}

func (l *Library) All() []*models.Book {
	return l.books
}

func (l *Library) findByTitle(title string) (*models.Book, error) {
	for _, b := range l.books {
		if b.Title == title {
			return b, nil
		}
	}
	return nil, fmt.Errorf("книга %q не найдена", title)
}

func (l *Library) Borrow(title string) error {
	b, err := l.findByTitle(title)
	if err != nil {
		return err
	}
	if b.Borrowed {
		return fmt.Errorf("книга %q уже на руках", title)
	}
	b.Borrowed = true
	return nil
}

func (l *Library) Return(title string) error {
	b, err := l.findByTitle(title)
	if err != nil {
		return err
	}
	if !b.Borrowed {
		return fmt.Errorf("книга %q не выдавалась", title)
	}
	b.Borrowed = false
	return nil
}
