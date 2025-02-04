package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := scanner.Text()

		if len(input) == 0 {
			continue
		}

		sliceOfInput := cleanInput(input)
		commands := getCommands()
		if command, ok := commands[sliceOfInput[0]]; ok {
			err := command.callback()
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
	}
}
