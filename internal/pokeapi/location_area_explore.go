package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *PokeApiClient) LocationAreaExplore(areaName string) (LocationArea, error) {
	fullUrl := baseUrl + "location-area/" + areaName

	data := LocationArea{}

	entry, found := c.cache.Get(fullUrl)
	if found {
		err := json.Unmarshal(entry, &data)
		if err != nil {
			return data, err
		}
	} else {
		res, err := c.httpClient.Get(fullUrl)
		if err != nil {
			return LocationArea{}, err
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return LocationArea{}, fmt.Errorf("response failed with status code: %s", res.Status)
		}

		bodySlice, err := io.ReadAll(res.Body)
		if err != nil {
			return LocationArea{}, err
		}

		err = json.Unmarshal(bodySlice, &data)
		if err != nil {
			return data, err
		}

		c.cache.Add(fullUrl, bodySlice)
	}

	return data, nil
}
