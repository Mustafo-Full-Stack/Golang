package main

import (
	"fmt"

	cs "struct-declaration/customstruct"
)

func main() {
	user := cs.User{
		ID:      "4234234234",
		Name:    "Ali",
		Age:     39,
		Address: "some address",
	}

	userAge := cs.GetAge(user)
	fmt.Println(userAge)

	userProfile := cs.GetProfile(user)
	fmt.Printf("%+v\n", userProfile)
}
