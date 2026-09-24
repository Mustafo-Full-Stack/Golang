package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Println("Сгенерированный ID:", id.String())

	// Каждый запуск - новый ID
	second := uuid.New()
	fmt.Println("Ещё один ID:      ", second.String())
	fmt.Println("Они разные:", id != second)
}
