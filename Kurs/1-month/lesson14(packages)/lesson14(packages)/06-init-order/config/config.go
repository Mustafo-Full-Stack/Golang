package config

import "fmt"

var (
	AppName string
	Version string
	Debug   bool
)

// init выполняется автоматически при загрузке пакета, ещё до main().
func init() {
	fmt.Println("пакет config: функция init() сработала ...")
	AppName = "MyApp"
	Version = "1.0.0"
	Debug = true
}

func Print() {
	fmt.Printf("%s v%s (debug=%t)\n", AppName, Version, Debug)
}
