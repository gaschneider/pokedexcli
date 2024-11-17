package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gaschneider/pokedexcli/internal/poketypes"
)

func (c *Client) GetPokemonCharacteristics(pokemonName string) (poketypes.PokemonCharacteristics, error) {
	url := baseURL + "/pokemon/" + pokemonName

	var pokemonCharacteristics poketypes.PokemonCharacteristics
	if cachedBody, exists := c.cache.Get(url); exists {
		result, err := parsePokemon(cachedBody)
		if err != nil {
			return pokemonCharacteristics, err
		}

		pokemonCharacteristics = result
	} else {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return pokemonCharacteristics, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return pokemonCharacteristics, fmt.Errorf("error looking for pokemon: %v - %v", pokemonName, err)
		}

		if res.StatusCode != 200 {
			return pokemonCharacteristics, fmt.Errorf("pokemon not found: %v", pokemonName)
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return pokemonCharacteristics, fmt.Errorf("error decoding response: %v", err)
		}

		result, err := parsePokemon(body)
		if err != nil {
			return pokemonCharacteristics, err
		}

		pokemonCharacteristics = result
		c.cache.Add(url, body)
	}

	return pokemonCharacteristics, nil
}

func parsePokemon(body []byte) (poketypes.PokemonCharacteristics, error) {
	var pokemonCharacteristics poketypes.PokemonCharacteristics
	if err := json.Unmarshal(body, &pokemonCharacteristics); err != nil {
		return poketypes.PokemonCharacteristics{}, fmt.Errorf("error decoding response: %v", err)
	}

	return pokemonCharacteristics, nil
}
