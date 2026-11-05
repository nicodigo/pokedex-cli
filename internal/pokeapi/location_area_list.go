package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *PokeApiClient) LocationAreaList(confUrl *string) (PokeApiListResponse, error) {
	data := PokeApiListResponse{}
	url := baseUrl + "location-area/"
	if confUrl != nil {
		url = *confUrl
	}

	entry, found := c.cache.Get(url)
	if found {
		err := json.Unmarshal(entry, &data)
		if err != nil {
			return data, err
		}
	} else {
		res, err := c.httpClient.Get(url)
		if err != nil {
			return PokeApiListResponse{}, err
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return PokeApiListResponse{}, fmt.Errorf("response failed with status code: %s", res.Status)
		}

		bodySlice, err := io.ReadAll(res.Body)
		if err != nil {
			return PokeApiListResponse{}, err
		}

		err = json.Unmarshal(bodySlice, &data)
		if err != nil {
			return data, err
		}

		c.cache.Add(url, bodySlice)
	}

	return data, nil
}
