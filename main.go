package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	err := callRepl()
	if err != nil {
		os.Exit(0)
	} else {
		fmt.Println(err)
	}
}
