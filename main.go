package main

import (
	"pokedexcli/inetrnal/pokeapi"
	"time"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: client,
	}
	repl(cfg)
}
