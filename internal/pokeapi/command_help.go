package pokeapi

import (
	"fmt"
)

func (api *PokeAPI) commandHelp(args ...string) error {
	fmt.Print("Welcome to the Pokedex!\n\nUsage:\n")
	for _, cmd := range api.GetCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
