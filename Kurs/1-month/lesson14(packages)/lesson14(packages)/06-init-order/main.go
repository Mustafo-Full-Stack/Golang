package main

import (
	"fmt"

	"myapp/config"
)

func init() {
	fmt.Println("пакет main init()")
}

func main() {
	fmt.Println("[main] main")
	config.Print()
}

// Порядок вывода:
//   [config] инициализация...
//   [main] init
//   [main] main
//   MyApp v1.0.0 (debug=true)
//
// Сначала init() импортированных пакетов (config), потом init() самого
// main, и только потом тело main(). Константы и переменные пакета
// готовы к этому моменту уже все.
