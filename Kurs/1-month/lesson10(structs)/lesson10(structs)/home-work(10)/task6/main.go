package main

type Movie struct {
	Title    string
	Duration int
	Rating   float64
}

func Films() Movie {
	info := Movie{
		Title:    "Маша и Медведь",
		Duration: 8,
		Rating:   4.9,
	}

	// info2 := Movie{
	// 	Title:    "Ну Погоди!",
	// 	Duration: 7,
	// 	Rating:   5.0,
	// }

	return info
}

func main() {

}

// func Films() Movie {
// 	fmt.Println("")
// }

// func Films() Movie {
// 	fmt.Println("")
// }

// func Films() Movie {
// 	fmt.Println("")
// }

// func Films() Movie {
// 	fmt.Println("")
// }

// func Films() Movie {
// 	fmt.Println("")
// }

// func Films() Movie {
// 	fmt.Println("")
// }
