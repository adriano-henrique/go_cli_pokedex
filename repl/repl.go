package repl

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommands struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	fmt.Println("Here is the list of available commands: ")
	fmt.Println()
	availableCommands := initializeCliCommands()
	for key, element := range availableCommands {
		fmt.Printf("%s : %s \n", key, element.description)
	}
	return nil
}

func commandExit() error {
	os.Exit(1)
	return nil
}

func initializeCliCommands() map[string]cliCommands {
	return map[string]cliCommands{
		"help": {
			name:        "help",
			description: "Shows the available REPL commands",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the pokedex",
			callback:    commandExit,
		},
	}
}

func SetupRepl() {
	fmt.Println("Welcome to the Pokedex REPL")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		for command, action := range initializeCliCommands() {
			if command == input {
				action.callback()
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input: ", err)
	}

	fmt.Println("Closing REPL")
}
