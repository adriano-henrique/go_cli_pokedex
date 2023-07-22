package repl

type cliCommands struct {
	name        string
	description string
	callback    func() error
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
