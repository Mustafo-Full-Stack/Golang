package main

import (
	"context"
	"crud-api/internal/models"
	userrepo "crud-api/internal/repository/user"
	"log"
)

func main() {
	ctx := context.Background()
	repo := userrepo.NewUserRepoInMemory()

	log.Println("main: вызываем repo.CreateUser")
	id, err := repo.CreateUser(ctx, &models.User{
		Name:         "Ali",
		Email:        "ali@example.com",
		PasswordHash: "plain-password",
	})
	if err != nil {
		log.Fatalf("main: repo.CreateUser: %v", err)
	}
	log.Printf("main: repo вернул id=%d", id)

	log.Println("main: вызываем repo.GetByID")
	user, err := repo.GetByID(ctx, id)
	if err != nil {
		log.Fatalf("main: repo.GetByID: %v", err)
	}
	log.Printf("main: repo вернул user=%+v", *user)

	log.Println("main: вызываем repo.GetAll")
	users, err := repo.GetAll(ctx)
	if err != nil {
		log.Fatalf("main: repo.GetAll: %v", err)
	}
	log.Printf("main: repo вернул %d пользователей", len(users))
}
