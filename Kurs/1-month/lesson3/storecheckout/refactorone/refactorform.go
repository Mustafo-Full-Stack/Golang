package main

import (
	"fmt"
	"strings"
)

/*
Условие:

Пользователь вводит:
цену товара;
количество;
есть ли скидочная карта.

Правила:
если есть скидочная карта → скидка 10%;
если сумма больше 1000 → доставка бесплатно;
иначе доставка 100 сомони.
*/

type Answer string

const (
	YES Answer = "ДА"
	NO  Answer = "НЕТ"
)

func main() {
	var (
		price      float64
		count      int
		cardAnswer string
		hasCard    bool
	)

	fmt.Print("Введите цену товара: ")
	fmt.Scan(&price)

	// Валидация цены
	if price <= 0 {
		fmt.Println("Ошибка: цена должна быть больше 0")
		return
	}

	fmt.Print("Введите количество: ")
	fmt.Scan(&count)

	// Валидация количества
	if count <=0 {
		fmt.Println("Ошибка: количество должно быть больше 0")
		return
	}

	fmt.Print("Есть скидочная карта? (да/нет): ")
	fmt.Scan(&cardAnswer)

	answer := Answer(strings.ToUpper(cardAnswer))

	switch answer {
	case YES:
		hasCard = true

	case NO:
		hasCard = false

	default:
		fmt.Println("Ошибка: введите да или нет")
		return
	}

	total := price * float64(count)

	fmt.Println("Стоимость без скидки:", total)

	if hasCard {
		discount := total * 0.10
		total -= discount

		fmt.Println("Применена скидка 10%")
	}

	var delivery float64

	if total > 1000 {
		fmt.Println("Доставка бесплатная")
	} else {
		delivery = 100
		fmt.Println("Стоимость доставки:", delivery)
	}

	finalPrice := total + delivery

	fmt.Println("Итого:", finalPrice)
}
