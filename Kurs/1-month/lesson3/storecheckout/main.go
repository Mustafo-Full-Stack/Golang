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
type Money float64

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

	fmt.Print("Введите количество: ")
	fmt.Scan(&count)

	fmt.Print("Есть скидочная карта? (да/нет): ")
	fmt.Scan(&cardAnswer)

	/*
		не очень красиво!

		switch strings.ToUpper(cardAnswer) {
			case "ДА":
				hasCard = true

			case "НЕТ":
				hasCard = false

			default:
				fmt.Println("Ошибка: введите да или нет")
				return
			}

	*/

	answer := Answer(strings.ToUpper(cardAnswer))
	//answer := strings.ToUpper(cardAnswer)

	switch answer {
	case YES:
		hasCard = true
	case NO:
		hasCard = false
	default:
		fmt.Println("Ошибка: введите да или нет")
		return
	}

	orderSum := price * float64(count)

	fmt.Println("Стоимость без скидки:", orderSum)

	if hasCard {
		discount := orderSum * 0.10
		orderSum -= discount
		// total = total - discount

		fmt.Println("Применена скидка 10%")
	}

	var delivery float64

	if orderSum > 1000 {
		fmt.Println("Доставка бесплатная")
	} else {
		delivery = 100
		fmt.Println("Стоимость доставки:", delivery)
	}

	finalSum := orderSum + delivery

	fmt.Println("Итого:", finalSum)
}
