package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	cleanedSlice := strings.Fields(strings.ToLower(text))
	return cleanedSlice
}

func (api *PokeApi) getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the program",
			callback:    api.commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    api.commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 names of Pokemon locations",
			callback:    api.commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 names of Pokemon locations",
			callback:    api.commandMapb,
		},
	}
}

func callRepl() error {
	scanner := bufio.NewScanner(os.Stdin)
	api := new(PokeApi)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}

		cmdName := text[0]

		cmd, ok := api.getCommands()[cmdName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		} else {
			err := cmd.callback()
			if err != nil {
				return err
			}
		}
	}
}
