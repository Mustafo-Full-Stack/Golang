package main

import (
	"fmt"
	"os"

	"demo-project/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "app stopped:", err)
		os.Exit(1)
	}
}
