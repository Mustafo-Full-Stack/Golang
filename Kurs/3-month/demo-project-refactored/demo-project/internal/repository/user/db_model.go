package user

import (
	"time"

	"demo-project/internal/models"
)

type dbModel struct {
	ID           int64     `db:"id"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	Name         string    `db:"name"`
	RoleID       int64     `db:"role_id"`
	RoleCode     string    `db:"role_code"`
	CreatedAt    time.Time `db:"created_at"`
}

func (m dbModel) toModel() models.User {
	return models.User{
		ID:           m.ID,
		Email:        m.Email,
		PasswordHash: m.PasswordHash,
		Name:         m.Name,
		RoleID:       m.RoleID,
		RoleCode:     m.RoleCode,
		CreatedAt:    m.CreatedAt,
	}
}
