package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gaschneider/pokedexcli/internal/poketypes"
)

func (c *Client) ListPokemonPerArea(areaToLookForPokemon string) (result poketypes.PokemonEncounterPerArea, err error) {
	url := baseURL + "/location-area/" + areaToLookForPokemon

	if cachedBody, exists := c.cache.Get(url); exists {
		result, err = parsePokemonPerArea(cachedBody)
		if err != nil {
			return result, err
		}

		return result, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return poketypes.PokemonEncounterPerArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return poketypes.PokemonEncounterPerArea{}, fmt.Errorf("error getting location areas: %v", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return poketypes.PokemonEncounterPerArea{}, fmt.Errorf("error decoding response: %v", err)
	}

	result, err = parsePokemonPerArea(body)
	if err != nil {
		return result, err
	}

	c.cache.Add(url, body)

	return result, nil
}

func parsePokemonPerArea(body []byte) (poketypes.PokemonEncounterPerArea, error) {
	var pokemonEncounterPerArea poketypes.PokemonEncounterPerArea
	if err := json.Unmarshal(body, &pokemonEncounterPerArea); err != nil {
		return poketypes.PokemonEncounterPerArea{}, fmt.Errorf("error decoding response: %v", err)
	}

	return pokemonEncounterPerArea, nil
}
