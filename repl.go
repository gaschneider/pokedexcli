package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gaschneider/pokedexcli/internal/pokeapi"
	"github.com/gaschneider/pokedexcli/internal/poketypes"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          map[string]poketypes.PokemonCharacteristics
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Explore the world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Explore the world backwards",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <area_name>",
			description: "Explore an area to encounter pokemons",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Try catching a pokemon!",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Check details on your caught pokemons",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Check the list of all your pokemons",
			callback:    commandPokedex,
		},
	}
}

func startREPL(cfg *config) {
	availableCommands := getCommands()
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		// Reads user input until user hit enter
		input := cleanInput(reader.Text())
		if len(input) == 0 {
			continue
		}

		commandName := input[0]
		params := []string{}
		if len(input) > 1 {
			params = input[1:]
		}

		if command, exists := availableCommands[commandName]; exists {
			if err := command.callback(cfg, params...); err != nil {
				fmt.Println(err)
			}
			continue
		}

		invalidCommand()
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
