package main

import (
	"fmt"
)

func commandPokedex(cfg *config, _ ...string) error {

	fmt.Println("Your Pokedex:")

	for k := range cfg.pokedex {
		fmt.Printf(" - %v\n", k)
	}

	fmt.Println()

	return nil
}
