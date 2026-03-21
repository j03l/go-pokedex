package pokeapi

import "fmt"

func (api *PokeAPI) commandPokedex(args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range api.CaughtPokemon {
		fmt.Printf("	- %s\n", pokemon.Name)
	}
	return nil
}
