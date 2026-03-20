package pokeapi

import "fmt"

func (api *LocationArea) commandExplore(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("A location name is required.")
	}

	url := "https://pokeapi.co/api/v2/location-area/" + args[0] + "/"

	if _, err := loadData(url, api); err != nil { // maybe cache later
		return fmt.Errorf("Failed to explore location: %w", err)
	}

	fmt.Printf("Exploring %s...\nFound Pokemon:\n", args[0])
	for _, enc := range api.PokemonEncounters {
		fmt.Printf("- %s\n", enc.Pokemon.Name)
	}
	return nil
}
