package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("explore command requires a location")
	}
	fmt.Printf("Exploring %s...:\n", args[0])
	pokemon, err := cfg.pokeapiClient.GetLocation(&args[0])
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for i := 0; i < len(pokemon.PokemonEncounters); i++ {
		fmt.Printf("- %s\n", pokemon.PokemonEncounters[i].Pokemon.Name)
	}
	return nil
}
