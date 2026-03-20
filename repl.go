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

func callRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	api := new(pokeapi.LocationArea) // just locations for now

	interval := time.Minute * 5
	api.Cache = pokecache.NewCache(interval)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		args := cleanInput(scanner.Text())
		if len(args) == 0 {
			continue
		}

		cmdName := args[0]
		cmd, ok := api.GetCommands()[cmdName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := cmd.Callback(args[1:]...)
		if err != nil {
			fmt.Println(err)
			continue // handle all errors inside repl
		}
	}
}
