package main

import (
	"fmt"
)

func (api *PokeApi) commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\n\nUsage:\n")
	for _, cmd := range api.getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
