package pokeapi

import (
	"fmt"
	"os"
)

func (api *LocationsAreaApi) commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
