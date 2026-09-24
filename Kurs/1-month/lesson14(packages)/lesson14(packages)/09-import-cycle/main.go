package main

import (
	"fmt"

	"myapp/models"
	"myapp/storage"
)

func main() {
	s := &storage.Storage{}
	s.Add(models.User{ID: 1, Name: "Али"})
	s.Add(models.User{ID: 2, Name: "Фарход"})

	for _, u := range s.All() {
		fmt.Printf("User#%d %s\n", u.ID, u.Name)
	}
}
