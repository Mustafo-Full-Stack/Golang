package users

import (
	"errors"
)

type User struct {
	Name  string
	Email string
}

func GetUser(name string, email string) (User, error) {
	user1 := User{Name: name, Email: email}
	if user1.Name == "" {
		return User{}, errors.New("имя не должно быть пустым")

	}

	return user1, nil
}

// func (u User) GetError() error {
// 	if User.Name == "" {
// 		return errors.New("имя не должен быть пустым!")
// 	}

// 	return nil
// }
