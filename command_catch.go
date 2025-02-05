package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("catch command requires a pokemon")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(&args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])

	if catchResult(float64(pokemon.BaseExperience)) {
		fmt.Printf("%s was caught!\n", args[0])
		cfg.pokemonCaught[args[0]] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", args[0])
	}
	return nil
}

func catchResult(val float64) bool {
	minVal := 39.0
	maxVal := 324.0
	if val == minVal {
		return rand.Float64()*100 <= 90.0
	}
	if val == maxVal {
		return rand.Float64()*100 <= 10.0
	}
	return rand.Float64()*100 <= 10.0+(90.0-10.0)*(val-minVal)/(maxVal-minVal)
}
