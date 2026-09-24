package main

import (
	"fmt"

	"myapp/internal/secret"
)

func main() {
	// Изнутри модуля myapp пакет internal/secret импортируется как обычно.
	fmt.Println("ключ:", secret.APIKey())

	// Но если бы кто-то снаружи (другой модуль, например otherapp)
	// попробовал сделать import "myapp/internal/secret", он бы получил:
	//   use of internal package myapp/internal/secret not allowed
	// Компилятор проверяет это по пути: любой пакет внутри internal/
	// виден только модулю, в котором лежит эта папка internal.
}
