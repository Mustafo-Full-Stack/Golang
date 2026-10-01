package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Employee struct {
	Name     string
	Position string
	Salary   string
}

func writeCSV(filename string, employees []Employee) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	w.Write([]string{"Имя", "Должность", "Зарплата"}) // заголовок
	for _, e := range employees {
		if err := w.Write([]string{e.Name, e.Position, e.Salary}); err != nil {
			return err
		}
	}
	return nil
}

func readCSV(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return csv.NewReader(f).ReadAll()
}

func main() {
	employees := []Employee{
		{"Али", "Разработчик", "5000"},
		{"Вера", "Дизайнер", "4000"},
		{"Тимур", "Тестировщик", "3500"},
	}
	if err := writeCSV("employees.csv", employees); err != nil {
		fmt.Println("Ошибка записи:", err)
		return
	}

	records, err := readCSV("employees.csv")
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}

	for i, rec := range records {
		fmt.Printf("%-10s %-15s %s\n", rec[0], rec[1], rec[2])
		if i == 0 {
			fmt.Println("-----------------------------------")
		}
	}
	fmt.Println("\nemployees.csv можно открыть в Excel/Google Sheets")
}
