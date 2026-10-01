package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrFileNotOpen = errors.New("[func appendLine]:не удалось открыть")
)

func appendLine(filename, text string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("[func appendLine]:не удалось открыть %s: %w", filename, err)
	}
	defer f.Close() // сразу после открытия - невозможно забыть

	if _, err := f.WriteString(text + "\n"); err != nil {
		return fmt.Errorf("ошибка записи: %w", err)
	}
	return nil
}

func main() {
	err := appendLine("log.txt", "Первая запись")
	if err != nil {
		if errors.Is(err, ErrFileNotOpen) {
			fmt.Println("cорри файл битый((")
		}
		fmt.Println(err)
	}
	//appendLine("log.txt", "Вторая запись")
	//appendLine("log.txt", "Третья запись")

	content, _ := os.ReadFile("log.txt")
	fmt.Println(string(content))
	fmt.Println("Запусти программу ещё раз - записи добавятся, а не перезапишутся")
}
