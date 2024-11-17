package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, _ ...string) error {
	fmt.Println()
	areas, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationsURL)
	if err != nil {
		fmt.Println(err)
		return err
	}

	cfg.nextLocationsURL = areas.Next
	cfg.prevLocationsURL = areas.Previous

	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}

	fmt.Println()

	return nil
}

func commandMapb(cfg *config, _ ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("can't explore backwards if no area was explored yet")
	}

	fmt.Println()
	areas, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationsURL)
	if err != nil {
		fmt.Println(err)
		return err
	}

	cfg.nextLocationsURL = areas.Next
	cfg.prevLocationsURL = areas.Previous

	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}

	fmt.Println()

	return nil
}
