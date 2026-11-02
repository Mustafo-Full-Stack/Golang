package models

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	minPasswordLen    = 8
	maxPasswordBytes  = 72
	maxHabitNameRunes = 100
)

type User struct {
	ID           int64
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}

type Habit struct {
	ID        int64
	UserID    int64
	Name      string
	CreatedAt time.Time
}

type HabitCheck struct {
	ID      int64
	HabitID int64
	Date    time.Time
}

type RegisterInput struct {
	Name     string
	Email    string
	Password string
}

func (in *RegisterInput) Validate() error {
	in.Name = strings.TrimSpace(in.Name)
	in.Email = strings.ToLower(strings.TrimSpace(in.Email))

	if in.Name == "" || in.Email == "" || in.Password == "" {
		return ErrInvalidInput
	}
	if len(in.Password) < minPasswordLen {
		return ErrPasswordTooShort
	}
	if len(in.Password) > maxPasswordBytes {
		return ErrPasswordTooLong
	}
	return nil
}

type LoginInput struct {
	Email    string
	Password string
}

func (in *LoginInput) Validate() error {
	in.Email = strings.ToLower(strings.TrimSpace(in.Email))
	if in.Email == "" || in.Password == "" {
		return ErrInvalidInput
	}
	return nil
}

type CreateHabitInput struct {
	UserID int64
	Name   string
}

func (in *CreateHabitInput) Validate() error {
	in.Name = strings.TrimSpace(in.Name)
	if in.Name == "" || utf8.RuneCountInString(in.Name) > maxHabitNameRunes {
		return ErrInvalidInput
	}
	return nil
}

var (
	ErrNotFound           = errors.New("not found")
	ErrInvalidInput       = errors.New("required fields are missing")
	ErrPasswordTooShort   = errors.New("password must be at least 8 characters")
	ErrPasswordTooLong    = errors.New("password must be at most 72 bytes")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUserExists         = errors.New("user already exists")
	ErrInvalidToken       = errors.New("invalid or missing token")
	ErrForbidden          = errors.New("access denied")
	ErrAlreadyChecked     = errors.New("habit already checked today")
)
