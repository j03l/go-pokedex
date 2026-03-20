package pokeapi

import (
	"fmt"
	"os"
)

func (api *LocationArea) commandExit(args ...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
