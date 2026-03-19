package main

import (
	"fmt"
	"os"
)

func main() {
	err := callRepl()
	if err != nil {
		fmt.Println(err)
	} else {
		os.Exit(0)
	}
}
