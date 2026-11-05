package main

import (
	"fmt"
)

func commandExplore(conf *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("location-area name missing")
	}

	locationArea, err := conf.client.LocationAreaExplore(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemon:")

	for _, pokemonEncounter := range locationArea.PokemonEncounters {
		pokemonName := pokemonEncounter.Pokemon.Name
		fmt.Println(" -", pokemonName)
	}

	return nil
}
