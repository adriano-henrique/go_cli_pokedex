package repl

type cliCommands struct {
	name        string
	description string
	callback    func(mapConfig *MapConfig) error
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
		"map": {
			name:        "map",
			description: "Gets the next location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Gets the previous location areas",
			callback:    commandMapb,
		},
	}
}
