package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *PokeApiClient) LocationAreaList(url string) (PokeApiListResponse, error) {
	res, err := c.httpClient.Get(url)
	if err != nil {
		return PokeApiListResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return PokeApiListResponse{}, fmt.Errorf("response failed with status code: %s", res.Status)
	}

	data := PokeApiListResponse{}
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return data, err
	}

	return data, nil
}
