package main

import (
	"fmt"
)

func commandMapf(conf *config, args []string) error {
	var fullUrl *string = nil

	if conf.locationNext != nil {
		fullUrl = conf.locationNext
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
