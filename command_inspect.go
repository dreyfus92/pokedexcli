package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return fmt.Errorf("You have not caught this pokemon yet")
	}
	
	fmt.Printf("%s\n", pokemon.Name)
	fmt.Printf("Base Experience: %v\n", pokemon.BaseExperience)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	for _,stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	for _,type_ := range pokemon.Types {
		fmt.Printf(" - %s\n", type_.Type.Name)
	}
	return nil
}
 