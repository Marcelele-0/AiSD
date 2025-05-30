package main

import (
	"fmt"
)

func main() {
	for {
		fmt.Println("🌳 BST Demo Program")
		fmt.Println("==================")
		fmt.Println("Wybierz opcję:")
		fmt.Println("1. Uruchom demo z małymi danymi (print)")
		fmt.Println("2. Uruchom testy wydajności (benchmark)")
		fmt.Println("3. Wyjście")
		fmt.Print("Twój wybór: ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Błąd odczytu. Spróbuj ponownie.")
			continue
		}

		switch choice {
		case 1:
			printMain()
		case 2:
			runBenchmarkMain()
		case 3:
			fmt.Println("Do widzenia!")
			return
		default:
			fmt.Println("Nieprawidłowy wybór. Spróbuj ponownie.")
		}
	}
}
