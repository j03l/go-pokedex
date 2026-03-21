package pokeapi

import "github.com/j03l/go-pokedex/internal/pokecache"

type PokeAPI struct {
	Cache         *pokecache.Cache   `json:"-"`
	CaughtPokemon map[string]Pokemon `json:"-"`
	Count         int                `json:"count"`
	Next          string             `json:"next"`
	Previous      string             `json:"previous"`
	Location      []struct {
		Name string `json:"name"`
	} `json:"results"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name           string `json:"name"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	BaseExperience int    `json:"base_experience"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}
