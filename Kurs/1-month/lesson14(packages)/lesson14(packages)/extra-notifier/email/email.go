package email

import (
	"fmt"

	"notifier/notifier"
)

// Sender реализует notifier.Notifier, но сам пакет notifier об этом
// даже не подозревает - в Go достаточно реализовать методы интерфейса,
// без явного "implements" как в других языках.
type Sender struct {
	Address string
}

func (s Sender) Send(message string) error {
	if message == "" {
		return notifier.ErrEmptyMessage
	}
	fmt.Printf("[email -> %s] %s\n", s.Address, message)
	return nil
}
