package main

import "fmt"

type Category struct {
	Name string
	Sub  []Category
}

func CountAll(c Category) int {
	count := 1 

	for _, sub := range c.Sub {
		count += CountAll(sub)
	}

	return count
}

func main() {
	root := Category{
		Name: "Товары",
		Sub: []Category{
			{
				Name: "Электроника",
				Sub: []Category{
					{Name: "Телефоны"},
					{Name: "Ноутбуки"},
				},
			},
			{
				Name: "Одежда",
				Sub: []Category{
					{Name: "Мужская"},
					{Name: "Женская"},
					{
						Name: "Детская",
						Sub: []Category{
							{Name: "Для мальчиков"},
							{Name: "Для девочек"},
						},
					},
				},
			},
		},
	}

	fmt.Println(CountAll(root))
}
