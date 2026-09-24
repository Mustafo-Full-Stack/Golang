package notifier

import "errors"

// ErrEmptyMessage - общая sentinel-ошибка. Любой пакет, который
// реализует Notifier, может её возвращать, а вызывающий код проверяет
// через errors.Is, не завязываясь на то, кто именно её вернул.
var ErrEmptyMessage = errors.New("сообщение не может быть пустым")

// Notifier - интерфейс из одного метода. Он живёт в отдельном пакете,
// а не внутри email или sms - так оба они зависят только от notifier,
// но не друг от друга. Циклического импорта тут в принципе не может
// возникнуть.
type Notifier interface {
	Send(message string) error
}
