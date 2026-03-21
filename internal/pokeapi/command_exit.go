package pokeapi

import (
	"fmt"
	"os"
)

func (api *PokeAPI) commandExit(args ...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
