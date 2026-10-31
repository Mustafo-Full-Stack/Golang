package models

import (
	"fmt"
	"net/netip"
	"strings"
	"time"
)

type DeviceType string

const (
	DeviceWeb    DeviceType = "web"
	DeviceMobile DeviceType = "mobile"
)

type Session struct {
	ID         string
	UserID     int64
	DeviceType DeviceType
	IP         netip.Addr
	UserAgent  string
	CreatedAt  time.Time
	ExpiresAt  time.Time
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type LoginInput struct {
	Email      string
	Password   string
	DeviceType DeviceType
	IP         netip.Addr
	UserAgent  string
}

// Пустой тип устройства — web; неизвестный — ошибка: молча подменять опечатку клиента нельзя.
func (in *LoginInput) Validate() error {
	in.Email = NormalizeEmail(in.Email)
	in.DeviceType = DeviceType(strings.ToLower(strings.TrimSpace(string(in.DeviceType))))
	if in.DeviceType == "" {
		in.DeviceType = DeviceWeb
	}

	if in.Email == "" || in.Password == "" {
		return fmt.Errorf("%w: email and password are required", ErrInvalidInput)
	}
	if in.DeviceType != DeviceWeb && in.DeviceType != DeviceMobile {
		return fmt.Errorf("%w: device_type must be web or mobile", ErrInvalidInput)
	}
	return nil
}
