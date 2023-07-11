package main

import (
	"time"

	"github.com/dreyfus92/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
	caughtPokemon map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}

	startRepl(&cfg)
}
