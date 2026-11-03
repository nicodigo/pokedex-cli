package main

import (
	"fmt"

	"github.com/nicodigo/pokedex-cli/internal/pokeapi"
)

func commandMapf(conf *config) error {
	fullUrl := pokeapi.BaseUrl + "location-area/"

	if conf.locationNext != nil {
		fullUrl = *conf.locationNext
	}

	location_areas, err := conf.client.LocationAreaList(fullUrl)
	if err != nil {
		return err
	}

	conf.locationNext = location_areas.Next
	conf.locationPrev = location_areas.Previous

	for _, area := range location_areas.Results {
		fmt.Println(area.Name)
	}

	return nil
}
