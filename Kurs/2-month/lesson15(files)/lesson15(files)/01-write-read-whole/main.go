package main

import (
	"fmt"
	"os"
)

func main() {
	// запись файла целиком
	//fileContentMsg := "Hello World"
	//if err := os.WriteFile("test.txt", []byte(fileContentMsg), 0644); err != nil {
	//	fmt.Println("Ошибка записи:", err)
	//	return
	//}
	//fmt.Println("Файл записан")
	//
	// чтение файла целиком
	content, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}
	fmt.Println("Содержимое:")
	//fmt.Println(content)
	fmt.Println(string(content))
	fmt.Println("Размер:", len(content), "байт")

	// os.WriteFile перезаписывает файл целиком - проверим
	os.WriteFile("test.txt", []byte("только эта строка"), 0644)
	content, _ = os.ReadFile("test.txt")
	fmt.Println("\nПосле повторной записи:")
	fmt.Println(string(content))

	// чтение несуществующего файла - современный способ проверки ошибки
	_, err = os.ReadFile("не_существует.txt")
	//if errors.Is(err, os.ErrNotExist) {
	//	fmt.Println("\nФайла нет, создадим значения по умолчанию")
	//}
	fmt.Println(err)
}
