package models

import (
	"fmt"
	"strings"
)

// User - модель пользователя.
type User struct {
	ID       int
	Name     string
	Email    string
	password string // приватное поле! видно только внутри пакета models
}

// NewUser - функция-конструктор с валидацией.
func New(id int, name, email string) (*User, error) {
	if name == "" {
		return nil, fmt.Errorf("имя не может быть пустым")
	}
	if !strings.Contains(email, "@") {
		return nil, fmt.Errorf("некорректный email: %s", email)
	}
	return &User{ID: id, Name: name, Email: email}, nil
}

// SetPassword - единственный способ задать пароль снаружи пакета.
func (u *User) SetPassword(p string) error {
	if len(p) < 8 {
		return fmt.Errorf("пароль слишком короткий")
	}
	u.password = p
	return nil
}

// CheckPassword проверяет пароль.
func (u *User) CheckPassword(p string) bool {
	return u.password == p
}

func (u User) String() string {
	return fmt.Sprintf("User#%d %s <%s>", u.ID, u.Name, u.Email)
}
