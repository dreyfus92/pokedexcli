package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

// Define the command functions
func commandHelp() error {
	fmt.Println("Help menu")
	return nil
}

func commandExit() error {
	fmt.Println("Exiting...")
	return nil
}


func startRepl() {
	newScanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print(" >")

		newScanner.Scan()
		text := newScanner.Text()

		fmt.Println("You entered:", text)
		
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
