package user

import (
	"crud-api/internal/models"
	"time"
)

type dbModel struct {
	ID           int64     `db:"id"`
	Name         string    `db:"name"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
}

func (m dbModel) toModel() models.User {
	return models.User{
		ID:           models.ID(m.ID),
		Name:         m.Name,
		Email:        m.Email,
		PasswordHash: m.PasswordHash,
		CreatedAt:    m.CreatedAt,
	}
}
