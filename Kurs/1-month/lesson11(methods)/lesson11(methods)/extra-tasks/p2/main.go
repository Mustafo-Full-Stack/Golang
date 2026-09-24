package main

import "fmt"

type User struct {
	ID   int
	Name string
}

// NewIDGenerator возвращает функцию,
// которая при каждом вызове выдает следующий ID.
func NewIDGenerator() func() int {
	id := 0

	return func() int {
		id++
		return id
	}
}

func main() {
	nextID := NewIDGenerator()

	users := []User{
		{
			ID:   nextID(),
			Name: "Ali",
		},
		{
			ID:   nextID(),
			Name: "Bob",
		},
		{
			ID:   nextID(),
			Name: "John",
		},
		{
			ID:   nextID(),
			Name: "Kate",
		},
		{
			ID:   nextID(),
			Name: "Sara",
		},
	}

	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s\n", user.ID, user.Name)
	}
}
