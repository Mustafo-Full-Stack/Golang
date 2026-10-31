package v1

import (
	"time"

	"demo-project/internal/models"
)

// userResponse — пользователь в ответе API. Пароль и role_id наружу не отдаём.
type userResponse struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func newUserResponse(u *models.User) userResponse {
	return userResponse{
		ID:        u.ID,
		Email:     u.Email,
		Name:      u.Name,
		Role:      u.RoleCode,
		CreatedAt: u.CreatedAt,
	}
}
