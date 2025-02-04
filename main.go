package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Pokedex > ")
		scanner.Scan()

		imput := scanner.Text()

		if len(imput) == 0 {
			continue
		}

		sliceOfImput := cleanInput(imput)

		fmt.Printf("Your command was: %s\n", sliceOfImput[0])
	}

}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	sliceOfText := strings.Fields(text)
	return sliceOfText
}
