package pokeapi

import "github.com/j03l/go-pokedex/internal/pokecache"

type LocationArea struct {
	Cache    *pokecache.Cache `json:"-"`
	Count    int              `json:"count"`
	Next     string           `json:"next"`
	Previous string           `json:"previous"`
	Location []struct {
		Name string `json:"name"`
	} `json:"results"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
