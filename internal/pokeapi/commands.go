package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type cliCommand struct {
	name        string
	description string
	Callback    func(args ...string) error
}

func (api *PokeAPI) GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the program",
			Callback:    api.commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    api.commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 names of Pokemon locations",
			Callback:    api.commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 names of Pokemon locations",
			Callback:    api.commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Lists all Pokemon in the specified area. Takes 1 name arg.",
			Callback:    api.commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon and add it to your Pokedex",
			Callback:    api.commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught Pokemon",
			Callback:    api.commandInspect,
		},
	}
}

func (api *PokeAPI) cacheCheck(url string) error {
	cacheHit, cached := api.Cache.Get(url)
	if cached {
		if err := json.Unmarshal(cacheHit, api); err != nil {
			return err
		}
	} else {
		data, err := loadData(url, api)
		if err != nil {
			return err
		}
		api.Cache.Add(url, data)
	}
	return nil
}

func loadData(url string, target any) ([]byte, error) {
	req, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Failed to request to PokeAPI: %w", err)
	}
	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Bad response or request not found")
	}

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read request body: %w", err)
	}

	if err := json.Unmarshal(res, target); err != nil {
		return nil, err
	}

	return res, nil
}
