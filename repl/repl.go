package repl

import (
	"bufio"
	"fmt"
	"os"
)

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
