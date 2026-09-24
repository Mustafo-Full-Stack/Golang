package main

import (
	"errors"
	"fmt"

	"notifier/email"
	"notifier/notifier"
	"notifier/sms"
)

func main() {
	// Срез интерфейсов notifier.Notifier, а элементы - конкретные типы
	// из двух совсем разных пакетов. main.go знает только про интерфейс,
	// а email и sms между собой вообще не знакомы.
	channels := []notifier.Notifier{
		email.Sender{Address: "ali@mail.tj"},
		sms.Sender{Phone: "+992 900 00 00 00"},
	}

	for _, ch := range channels {
		if err := ch.Send("Занятие начинается в 18:00"); err != nil {
			fmt.Println("ошибка отправки:", err)
		}
	}

	// Пустое сообщение - оба пакета возвращают одну и ту же
	// sentinel-ошибку notifier.ErrEmptyMessage, хотя реализованы отдельно.
	for _, ch := range channels {
		err := ch.Send("")
		if errors.Is(err, notifier.ErrEmptyMessage) {
			fmt.Printf("%T: сообщение пустое, отправка отменена\n", ch)
		}
	}
}
