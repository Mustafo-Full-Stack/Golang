package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task2/modules"
)

func GetClient(id int) (modules.User, error) {
	resp, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", id))
	if err != nil {
		fmt.Println("Ошибка:", err)
		return modules.User{}, err
	}

	defer resp.Body.Close()

	var user modules.User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return modules.User{}, err
	}

	return user, nil

}
