package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/j03l/go-pokedex/internal/pokecache"
)

type LocationsAreaApi struct {
	Count     int    `json:"count"`
	Next      string `json:"next"`
	Previous  string `json:"previous"`
	Locations []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
	Cache *pokecache.Cache `json:"-"`
}

func (api *LocationsAreaApi) commandMap() error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if api.Next != "" {
		url = api.Next
	}
	// url is cache key
	data, check := api.Cache.Get(url)
	if check {
		if err := api.printData(data); err != nil {
			return err
		}
	} else {
		return api.getMap(url)
	}
	return nil
}

func (api *LocationsAreaApi) getMap(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Failed to request to PokeAPI: %w", err)
	}
	defer res.Body.Close()

	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Failed to read request body: %w", err)
	}

	if err := api.printData(jsonData); err != nil {
		return err
	}
	api.Cache.Add(url, jsonData) // save for later

	return nil
}

func (api *LocationsAreaApi) printData(jsonData []byte) error {
	if err := json.Unmarshal(jsonData, &api); err != nil {
		return fmt.Errorf("Error unmarshalling JSON: %v", err)
	}
	for _, loc := range api.Locations {
		fmt.Println(loc.Name)
	}
	return nil
}
