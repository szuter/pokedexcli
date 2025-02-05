package main

import (
	"pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	client := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: client,
	}
	repl(cfg)
}
