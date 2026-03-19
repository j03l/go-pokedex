package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (api *PokeApi) commandMap() error {
	url := "https://pokeapi.co/api/v2/location/"
	if api.Next != "" {
		url = api.Next
	}
	return api.getMap(url)
}

func (api *PokeApi) getMap(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Failed to request to PokeAPI: %w", err)
	}
	defer res.Body.Close()

	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Failed to read request body: %w", err)
	}

	if err := json.Unmarshal(jsonData, &api); err != nil {
		return fmt.Errorf("Error unmarshalling JSON: %v", err)
	}
	for _, loc := range api.Locations {
		fmt.Println(loc.Name)
	}
	return nil
}
