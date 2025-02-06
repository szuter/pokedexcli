package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("inspect requires a pokemon name")
	}
	pokemonName := args[0]
	pokemon, ok := cfg.pokemonCaught[pokemonName]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Println("  -", stat.Stat.Name, ":", stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Println("  -", t.Type.Name)
	}
	return nil
}
