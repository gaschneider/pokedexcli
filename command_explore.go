package main

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func commandExplore(cfg *config, params ...string) error {
	if len(params) != 1 {
		return fmt.Errorf("explore command requires an area name")
	}

	areaName := params[0]
	fmt.Printf("Exploring %v\n", areaName)
	pokemonsPerArea, err := cfg.pokeapiClient.ListPokemonPerArea(areaName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	caser := cases.Title(language.English)
	for _, encounter := range pokemonsPerArea.PokemonEncounters {
		fmt.Printf("- %v\n", caser.String(encounter.Pokemon.Name))
	}

	fmt.Println()

	return nil
}
