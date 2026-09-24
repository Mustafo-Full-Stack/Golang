package main

import "fmt"

// Самый жизненный пример урока: почему интерфейсы реально важны.
// Notifier описывает "умение отправить сообщение" - неважно, каким способом.
type Notifier interface {
	Send(message string) error
}

type EmailNotifier struct {
	Address string
}

func (e EmailNotifier) Send(msg string) error {
	fmt.Printf("[EMAIL → %s] %s\n", e.Address, msg)
	return nil
}

type SMSNotifier struct {
	Phone string
}

func (s SMSNotifier) Send(msg string) error {
	if len(msg) > 70 {
		return fmt.Errorf("SMS слишком длинная: %d символов", len(msg))
	}
	fmt.Printf("[SMS → %s] %s\n", s.Phone, msg)
	return nil
}

type TelegramNotifier struct {
	ChatID int
}

func (t TelegramNotifier) Send(msg string) error {
	fmt.Printf("[TELEGRAM → %d] %s\n", t.ChatID, msg)
	return nil
}

// notifyAll не знает и не должна знать про Email, SMS или Telegram.
// Она знает только про Notifier. Это и есть развязка (decoupling).
func notifyAll(notifiers []Notifier, message string) {
	for _, n := range notifiers {
		if err := n.Send(message); err != nil {
			fmt.Println("  Ошибка отправки:", err)
		}
	}
}

func main() {
	notifiers := []Notifier{
		EmailNotifier{Address: "ali@mail.tj"},
		SMSNotifier{Phone: "+992900000000"},
		TelegramNotifier{ChatID: 123456},
	}

	notifyAll(notifiers, "Ваш заказ готов!")

	fmt.Println()
	notifyAll(notifiers, "Очень длинное сообщение, которое точно не поместится в одну SMS, потому что превышает лимит в семьдесят символов")

	// Захотим добавить push-уведомления - создадим PushNotifier с методом Send
	// и допишем его в срез. notifyAll не изменится ни на строку. Вот зачем
	// на практике нужны интерфейсы: код, зависящий от Notifier, никогда
	// не приходится трогать при появлении нового способа уведомить пользователя.
}
