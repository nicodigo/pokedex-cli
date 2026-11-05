package main

import (
	"fmt"
)

func commandMapb(conf *config, args []string) error {

	if conf.locationPrev == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	fullUrl := conf.locationPrev

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
