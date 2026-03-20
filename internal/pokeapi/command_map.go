package pokeapi

import (
	"fmt"
)

func (api *LocationArea) commandMap(args ...string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if api.Next != "" {
		url = api.Next
	}

	err := api.cacheCheck(url)
	if err != nil {
		return err
	}
	for _, location := range api.Location {
		fmt.Println(location.Name)
	}
	return nil
}
