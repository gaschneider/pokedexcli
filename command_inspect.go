package main

import (
	"fmt"
)

func commandInspect(cfg *config, params ...string) error {
	if len(params) != 1 {
		return fmt.Errorf("inspect command requires a pokemon name")
	}

	pokemonName := params[0]
	characteristics, exists := cfg.pokedex[pokemonName]
	if !exists {
		return fmt.Errorf("no information to display about pokemon %v, try catching it first", pokemonName)
	}

	fmt.Printf("Name: %v\n", pokemonName)
	fmt.Printf("Height: %v\n", characteristics.Height)
	fmt.Printf("Weight: %v\n", characteristics.Weight)
	fmt.Println("Stats:")
	for _, stat := range characteristics.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")

	for _, t := range characteristics.Types {
		fmt.Printf("  - %v\n", t.Type.Name)
	}

	fmt.Println()

	return nil
}
