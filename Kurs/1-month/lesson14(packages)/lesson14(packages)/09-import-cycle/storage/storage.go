package storage

import "myapp/models"

// storage имеет право импортировать models - зависимость идёт в одну
// сторону: "хранилище знает про модель", а не наоборот.
type Storage struct {
	users []models.User
}

func (s *Storage) Add(u models.User) {
	s.users = append(s.users, u)
}

func (s *Storage) All() []models.User {
	return s.users
}
