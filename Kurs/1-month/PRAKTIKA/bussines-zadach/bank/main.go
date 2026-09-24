package main

import "fmt"

type USER struct {
	name     string
	balance  float64
	staus    bool
	withdraw float64
	email    string
}

var exit string

func GetName() USER {
	var UserName USER

	fmt.Println("=============")
	fmt.Println("Введите имя :")
	fmt.Println("=============")
	fmt.Scanln(&UserName.name)

	return UserName
}

func GetEmail() USER {
	var userEmail USER

	fmt.Println("=============")
	fmt.Println("Введите email:")
	fmt.Println("=============")

	fmt.Scanln(&userEmail.email)

	return userEmail
}

func GetBalance() USER {
	var UserBalance USER

	fmt.Println("=============")
	fmt.Println("Введите баланс :")
	fmt.Println("=============")

	fmt.Scanln(&UserBalance.balance)

	return UserBalance
}

func GetStatus() USER {
	var userStatus USER

	fmt.Println("=============")
	fmt.Println("Статус :")
	fmt.Println("=============")

	fmt.Scanln(&userStatus.staus)

	return userStatus
}

func GetWithdraw() USER {
	var userWinthdraw USER

	fmt.Println("=============")
	fmt.Println("Введите сумму который хотите снят:")
	fmt.Println("=============")

	fmt.Scanln(&userWinthdraw.withdraw)

	return userWinthdraw
}

func main() {

	UsrName := GetName()

	UsrEmail := GetEmail()

	UsrBalance := GetBalance()

	UsrStatus := GetStatus()

	UsrWithdraw := GetWithdraw()

	fmt.Println("Имя:", UsrName.name)
	fmt.Println("Почта:", UsrEmail.email)

	//fmt.Println("Баланс:", UsrBalance.balance, "Сомони"
	if UsrBalance.balance < 0 {
		fmt.Println("ОШИБКА ! Баланс НЕ должен быть отрицательным")
	} else if UsrBalance.balance >= 1000 {
		fmt.Println("у вас больше денег чем 1000, БАЛАНС:", UsrBalance.balance, "Сомони")
	} else {
		fmt.Println("Баланс:", UsrBalance.balance, "Сомони")
	}

	//fmt.Println("Снято от баланса: -", UsrWithdraw.withdraw, "Сомони")
	if UsrWithdraw.withdraw > UsrBalance.balance {
		fmt.Println("ОШИБКА ! Нельзя снять больше денег, чем есть на балансе !")
	} else {
		fmt.Println("Снято от баланса: -", UsrWithdraw.withdraw, "Сомони")
	}

	fmt.Println("Ваш новый Баланс:", UsrBalance.balance-UsrWithdraw.withdraw, "Сомони")

	if UsrStatus.staus == true {
		fmt.Println("Статус: ОНЛАЙН")
	} else {
		fmt.Println("Статус: ОФЛАЙН")
	}
}
