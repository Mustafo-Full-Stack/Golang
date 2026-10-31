package models

import "time"

type ID int64

type User struct {
	ID           ID
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}
