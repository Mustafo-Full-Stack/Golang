package v1

import (
	"time"

	"demo-project/internal/models"
)

type sessionResponse struct {
	ID         string    `json:"id"`
	DeviceType string    `json:"device_type"`
	IP         string    `json:"ip,omitempty"`
	UserAgent  string    `json:"user_agent,omitempty"`
	Current    bool      `json:"is_current"`
	CreatedAt  time.Time `json:"created_at"`
}

func newSessionResponse(s *models.Session, currentID string) sessionResponse {
	resp := sessionResponse{
		ID:         s.ID,
		DeviceType: string(s.DeviceType),
		UserAgent:  s.UserAgent,
		Current:    s.ID == currentID,
		CreatedAt:  s.CreatedAt,
	}
	if s.IP.IsValid() {
		resp.IP = s.IP.String()
	}
	return resp
}

type sessionListResponse struct {
	Items []sessionResponse `json:"items"`
}

func newSessionListResponse(sessions []models.Session, currentID string) sessionListResponse {
	items := make([]sessionResponse, 0, len(sessions))
	for i := range sessions {
		items = append(items, newSessionResponse(&sessions[i], currentID))
	}
	return sessionListResponse{Items: items}
}
