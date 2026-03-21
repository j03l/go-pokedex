package pokeapi

import "fmt"

// Name: pidgey
// Height: 3
// Weight: 18
// Stats:
//
//	-hp: 40
//	-attack: 45
//	-defense: 40
//	-special-attack: 35
//	-special-defense: 35
//	-speed: 56
//
// Types:
//   - normal
//   - flying

func (api *PokeAPI) commandInspect(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("A Pokemon name is required.")
	}
	pokemon, ok := api.CaughtPokemon[args[0]]
	if !ok {
		return fmt.Errorf("You haven't caught that Pokemon.")
	}

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", pokemon.Name, pokemon.Weight, pokemon.Height)
	printPokemonData(pokemon) // stats & types
	return nil
}

func printPokemonData(pokemon Pokemon) {
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("	-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, types := range pokemon.Types {
		fmt.Printf("	- %s\n", types.Type.Name)
	}
}
