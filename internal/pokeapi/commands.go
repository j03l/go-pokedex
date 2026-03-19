package pokeapi

type cliCommand struct {
	name        string
	description string
	Callback    func() error
}

func (api *LocationsAreaApi) GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the program",
			Callback:    api.commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    api.commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 names of Pokemon locations",
			Callback:    api.commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 names of Pokemon locations",
			Callback:    api.commandMapb,
		},
	}
}
