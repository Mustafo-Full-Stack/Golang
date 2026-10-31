package models

import (
	"fmt"
	"net/mail"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	minPasswordLen   = 8
	maxPasswordBytes = 72 // bcrypt учитывает только первые 72 байта пароля
	maxNameRunes     = 100
)

type User struct {
	ID           int64
	Email        string
	PasswordHash string
	Name         string
	RoleID       int64
	RoleCode     string // например "user": нужен только для ответа API
	CreatedAt    time.Time
}

// Роли во входе нет: зарегистрироваться можно только как user.
type RegisterInput struct {
	Email    string
	Password string
	Name     string
}

// Validate сначала приводит поля к единому виду, потом проверяет их.
// Pointer receiver — потому что метод меняет поля. Так устроены все XxxInput.
func (in *RegisterInput) Validate() error {
	in.Email = NormalizeEmail(in.Email)
	in.Name = strings.TrimSpace(in.Name)

	if addr, err := mail.ParseAddress(in.Email); err != nil || addr.Address != in.Email {
		return fmt.Errorf("%w: invalid email", ErrInvalidInput)
	}
	if n := utf8.RuneCountInString(in.Name); n == 0 || n > maxNameRunes {
		return fmt.Errorf("%w: name must be 1 to %d characters", ErrInvalidInput, maxNameRunes)
	}
	// Postgres не принимает NUL в text: без проверки это был бы 500.
	if strings.ContainsRune(in.Name, '\x00') {
		return fmt.Errorf("%w: name must not contain NUL", ErrInvalidInput)
	}
	if utf8.RuneCountInString(in.Password) < minPasswordLen {
		return fmt.Errorf("%w: password must be at least %d characters", ErrInvalidInput, minPasswordLen)
	}
	if len(in.Password) > maxPasswordBytes {
		return fmt.Errorf("%w: password must be at most %d bytes", ErrInvalidInput, maxPasswordBytes)
	}
	return nil
}

func NormalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}
