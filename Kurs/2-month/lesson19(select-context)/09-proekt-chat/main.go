package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Итоговый проект недели: горутины + каналы + WaitGroup + select + таймаут.

type Message struct {
	From string
	Text string
	Time time.Time
}

func user(name string, messages chan<- Message, wg *sync.WaitGroup, count int) {
	defer wg.Done()
	texts := []string{"Привет!", "Как дела?", "Кто-нибудь online?", "Го в Go", "До связи"}

	for i := 0; i < count; i++ {
		time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond)
		messages <- Message{
			From: name,
			Text: texts[rand.Intn(len(texts))],
			Time: time.Now(),
		}
	}
}

func main() {
	messages := make(chan Message)
	var wg sync.WaitGroup

	users := []string{"Али", "Вера", "Тимур", "Нигина"}
	for _, u := range users {
		wg.Add(1)
		go user(u, messages, &wg, 3)
	}

	// закрываем канал, когда все написали
	go func() {
		wg.Wait()
		close(messages)
	}()

	timeout := time.After(5 * time.Second)
	count := 0

	fmt.Println("=== ЧАТ ЗАПУЩЕН ===")
	for {
		select {
		case msg, ok := <-messages:
			if !ok {
				fmt.Printf("\n=== Чат завершён. Сообщений: %d ===\n", count)
				return
			}
			count++
			fmt.Printf("[%s] %-8s: %s\n", msg.Time.Format("15:04:05"), msg.From, msg.Text)

		case <-timeout:
			fmt.Println("\n=== Таймаут чата ===")
			return
		}
	}
}
