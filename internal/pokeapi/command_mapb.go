package pokeapi

import "fmt"

func (api *LocationArea) commandMapb(args ...string) error {
	url := "https://pokeapi.co/api/v2/location/"
	if api.Previous != "" {
		url = api.Previous
	} else {
		return fmt.Errorf("you're on the first page")
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
