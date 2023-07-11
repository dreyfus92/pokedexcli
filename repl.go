package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	newScanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		newScanner.Scan()

		words := cleanInput(newScanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args:= []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}

}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next page of locations areas",
			callback:    mapForwardLoc,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous page of locations areas",
			callback:    mapBackwardsLoc,
		},
		"explore": {
			name:        "explore",
			description: "Displays the pokemon in a location area",
			callback:	callbackExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught pokemon",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays the pokemon in the pokedex",
			callback:    callbackPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
