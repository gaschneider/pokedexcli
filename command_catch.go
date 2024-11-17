package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, params ...string) error {
	if len(params) != 1 {
		return fmt.Errorf("catch command requires a pokemon name")
	}

	pokemonName := params[0]
	pokemonCharacteristics, err := cfg.pokeapiClient.GetPokemonCharacteristics(pokemonName)
	if err != nil {
		return err
	}

	caught := decreasingProbability(pokemonCharacteristics.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)
	time.Sleep(5 * time.Millisecond)
	if caught {
		fmt.Printf("%v was caught!\n", pokemonName)
		cfg.pokedex[pokemonName] = pokemonCharacteristics
	} else {
		fmt.Printf("%v escaped!\n", pokemonName)
	}

	fmt.Println()

	return nil
}

func decreasingProbability(input int) bool {
	// Ensure input is non-negative
	if input < 0 {
		input = 0
	}

	// Calculate probability between 0 and 1
	// Using exponential decay: e^(-k*x) where k is calculated to match our requirements
	// To get ~87% at x=10 and ~35% at x=160, k ≈ 0.0065
	k := 0.0065
	probability := math.Exp(-k * float64(input))

	// Generate random number between 0 and 1
	randomValue := rand.Float64()

	// Return true if random value is less than probability
	return randomValue < probability
}
