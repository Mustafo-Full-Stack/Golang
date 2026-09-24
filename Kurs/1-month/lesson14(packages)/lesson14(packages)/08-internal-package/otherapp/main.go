package main

import (
	"fmt"

	"myapp/internal/secret"
)

func main() {
	fmt.Println(secret.APIKey())
}
