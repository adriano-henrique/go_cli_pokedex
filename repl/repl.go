package repl

import (
	"bufio"
	"fmt"
	"os"
)

type MapConfig struct {
	Next     *string
	Previous *string
}

func SetupRepl() {
	fmt.Println("Welcome to the Pokedex REPL")
	scanner := bufio.NewScanner(os.Stdin)
	initialUrl := "https://pokeapi.co/api/v2/location-area/"

	mapConfig := MapConfig{
		Next:     &initialUrl,
		Previous: nil,
	}

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		for command, action := range initializeCliCommands() {
			if command == input {
				action.callback(&mapConfig)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input: ", err)
	}

	fmt.Println("Closing REPL")
}
