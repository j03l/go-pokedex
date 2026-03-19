package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/j03l/go-pokedex/internal/pokeapi"
	"github.com/j03l/go-pokedex/internal/pokecache"
)

func cleanInput(text string) []string {
	cleanedSlice := strings.Fields(strings.ToLower(text))
	return cleanedSlice
}

func callRepl() error {
	scanner := bufio.NewScanner(os.Stdin)
	api := new(pokeapi.LocationsAreaApi) // just locations for now

	interval := time.Second * 10
	api.Cache = pokecache.NewCache(interval)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}

		cmdName := text[0]

		cmd, ok := api.GetCommands()[cmdName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		} else {
			err := cmd.Callback()
			if err != nil {
				return err
			}
		}
	}
}
