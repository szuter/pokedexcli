package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
	pokemonCaught map[string]pokeapi.Pokemon
}

func repl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := scanner.Text()

		if len(input) == 0 {
			continue
		}

		sliceOfInput := cleanInput(input)
		args := []string{}
		if len(sliceOfInput) > 1 {
			args = sliceOfInput[1:]
		}
		commands := getCommands()
		if command, ok := commands[sliceOfInput[0]]; ok {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command:", sliceOfInput[0])
		}
	}

}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	sliceOfText := strings.Fields(text)
	return sliceOfText
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Shows the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location name>",
			description: "Shows the pokemon in a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon name>",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
	}
}
