package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	newScanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" >")

		newScanner.Scan()
		text := newScanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		command := cleaned[0]

		switch command {
		case "exit":
			os.Exit(0)
		case "help":
			fmt.Println("Welcome to the Pokedex!")
			fmt.Println("=====================================")
			fmt.Println("Available commands:")
			fmt.Println("exit - exit the program")
			fmt.Println("help - prints the available commands")
		default:
			fmt.Println("Unknown command:", command)
		}

	}

}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
