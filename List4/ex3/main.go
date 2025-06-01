package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("🌳 RBBST Demo Program")
		fmt.Println("==================")
		fmt.Println("Wybierz opcję:")
		fmt.Println("1. Uruchom demo z małymi danymi (print)")
		fmt.Println("2. Uruchom testy wydajności (benchmark)")
		fmt.Println("3. Wyjście")
		fmt.Print("Twój wybór: ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Printf("Błąd odczytu: %v\n", err)
			}
			// EOF lub koniec pipe - kończymy program
			return
		}

		input := strings.TrimSpace(scanner.Text())
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Nieprawidłowy wybór. Spróbuj ponownie.")
			continue
		}

		switch choice {
		case 1:
			printMain()
			return // Kończymy po demo
		case 2:
			runBenchmarkMain()
			return // Kończymy po benchmark
		case 3:
			fmt.Println("Do widzenia!")
			return
		default:
			fmt.Println("Nieprawidłowy wybór. Spróbuj ponownie.")
		}
	}
}
