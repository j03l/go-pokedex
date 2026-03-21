package pokeapi

import (
	"fmt"
	"math/rand"
)

func (api *PokeAPI) commandCatch(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("A Pokemon name is required.")
	}

	var caughtPokemon Pokemon
	url := "https://pokeapi.co/api/v2/pokemon/" + args[0] + "/"
	if _, err := loadData(url, &caughtPokemon); err != nil {
		return fmt.Errorf("Failed to find Pokemon: %w", err)
	}

	namedPokemon := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", namedPokemon)

	rng := rand.Intn(caughtPokemon.BaseExperience) // base xp; arceus is 324, caterpie is 39
	if rng > 50 {                                  // TODO: pokeball will modify this
		fmt.Printf("%s escaped!\n", namedPokemon)
	} else {
		api.CaughtPokemon[namedPokemon] = caughtPokemon // saved to Pokedex
		fmt.Printf("%s was caught!\n", namedPokemon)
	}
	return nil
}
