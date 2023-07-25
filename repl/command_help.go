package repl

import "fmt"

func commandHelp(mapConfig *MapConfig) error {
	fmt.Println("Here is the list of available commands: ")
	fmt.Println()
	availableCommands := initializeCliCommands()
	for key, element := range availableCommands {
		fmt.Printf("%s : %s \n", key, element.description)
	}
	return nil
}
