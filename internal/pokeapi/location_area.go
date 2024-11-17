package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gaschneider/pokedexcli/internal/poketypes"
)

func (c *Client) ListLocationAreas(pageURL *string) (result poketypes.LocationAreas, err error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if cachedBody, exists := c.cache.Get(url); exists {
		result, err = parseAreas(cachedBody)
		if err != nil {
			return result, err
		}

		return result, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return poketypes.LocationAreas{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return poketypes.LocationAreas{}, fmt.Errorf("error getting location areas: %v", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return poketypes.LocationAreas{}, fmt.Errorf("error decoding response: %v", err)
	}

	result, err = parseAreas(body)
	if err != nil {
		return result, err
	}

	c.cache.Add(url, body)

	return result, nil
}

func parseAreas(body []byte) (poketypes.LocationAreas, error) {
	var locationAreas poketypes.LocationAreas
	if err := json.Unmarshal(body, &locationAreas); err != nil {
		return poketypes.LocationAreas{}, fmt.Errorf("error decoding response: %v", err)
	}

	return locationAreas, nil
}
