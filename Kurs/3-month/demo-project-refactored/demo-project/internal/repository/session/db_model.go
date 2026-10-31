package session

import (
	"net/netip"
	"time"

	"demo-project/internal/models"
)

type dbModel struct {
	ID         string      `db:"id"`
	UserID     int64       `db:"user_id"`
	DeviceType string      `db:"device_type"`
	IP         *netip.Addr `db:"ip"`
	UserAgent  string      `db:"user_agent"`
	CreatedAt  time.Time   `db:"created_at"`
	ExpiresAt  time.Time   `db:"expires_at"`
}

func (m dbModel) toModel() models.Session {
	s := models.Session{
		ID:         m.ID,
		UserID:     m.UserID,
		DeviceType: models.DeviceType(m.DeviceType),
		UserAgent:  m.UserAgent,
		CreatedAt:  m.CreatedAt,
		ExpiresAt:  m.ExpiresAt,
	}
	if m.IP != nil {
		s.IP = *m.IP
	}
	return s
}
