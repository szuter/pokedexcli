package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your pokedex:")
	for _, pokemon := range cfg.pokemonCaught {
		fmt.Println(" -", pokemon.Name)
	}
	return nil
}
