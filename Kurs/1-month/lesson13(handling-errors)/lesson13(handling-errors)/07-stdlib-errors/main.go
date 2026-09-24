package main

import (
	"fmt"
	"strconv"
)

func parseAge(s string) (int, error) {
	age, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("не удалось разобрать возраст %q: %w", s, err)
	}
	if age < 0 {
		return 0, fmt.Errorf("отрицательный возраст: %d", age)
	}
	return age, nil
}

func main() {
	for _, s := range []string{"25", "abc", "-5"} {
		age, err := parseAge(s)
		if err != nil {
			fmt.Println("✗", err)
			continue
		}
		fmt.Println("✓ Возраст:", age)
	}
}
