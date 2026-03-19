package pokeapi

import "fmt"

func (api *LocationsAreaApi) commandMapb() error {
	url := "https://pokeapi.co/api/v2/location/"
	if api.Previous != "" {
		url = api.Previous
	} else {
		fmt.Println("you're on the first page")
		return nil
	}
	return api.getMap(url)
}
