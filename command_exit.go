package main

import (
	"fmt"
	"os"
)

func (api *PokeApi) commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
