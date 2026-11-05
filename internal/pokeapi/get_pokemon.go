package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *PokeApiClient) GetPokemon(pokemonName string) (Pokemon, error) {
	fullUrl := baseUrl + "pokemon/" + pokemonName
	var pokemon Pokemon

	if data, found := c.cache.Get(fullUrl); found {
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return pokemon, err
		}
		return pokemon, nil
	}

	res, err := c.httpClient.Get(fullUrl)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("response failed with status code: %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return pokemon, err
	}

	c.cache.Add(fullUrl, data)
	return pokemon, nil
}
