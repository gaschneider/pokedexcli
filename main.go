package main

import (
	"time"

	"github.com/gaschneider/pokedexcli/internal/pokeapi"
	"github.com/gaschneider/pokedexcli/internal/poketypes"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)

	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       make(map[string]poketypes.PokemonCharacteristics),
	}
	startREPL(cfg)
}
