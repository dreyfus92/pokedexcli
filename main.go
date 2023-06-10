package main

import (
	"bufio"
	"fmt"
	"os"
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


func main() {

	// Create a map of commands
	commands := map[string]cliCommand{
		"help": {
			name: "help",
			description: "Prints the help menu",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandHelp,
		},
	}

	// Print the welcome message
	fmt.Printf("Welcome to the Pokedex!\n\n")

	// Print command list
	fmt.Println("Commands:")
	fmt.Println("==============================")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	fmt.Println("==============================")

	fmt.Println("Enter a command:")

	// Get user input 
	for {
		newScanner := bufio.NewScanner(os.Stdin)
		newScanner.Scan()
		userInput := newScanner.Text()
		fmt.Println("You entered: ")
		fmt.Println(userInput)

		// Check if the command exists
		if command, ok := commands[userInput]; ok {
			// If it does, execute the command
			command.callback()
		} else {
			// If it doesn't, print an error message
			fmt.Printf("Unknown command: %s\n", userInput)
		}
	}
}