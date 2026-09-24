package sms

import (
	"fmt"

	"notifier/notifier"
)

type Sender struct {
	Phone string
}

func (s Sender) Send(message string) error {
	if message == "" {
		return notifier.ErrEmptyMessage
	}
	fmt.Printf("[sms -> %s] %s\n", s.Phone, message)
	return nil
}
