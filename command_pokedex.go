package main

import "fmt"

func commandPokedex(conf *config, args []string) error {
	fmt.Println("Your Pokedex:")
	if len(conf.pokemonRegistry) == 0 {
		fmt.Println("  is empty, use the 'catch' command")
		fmt.Println("  to catch a pokemon and add it to the pokedex")
		fmt.Println()
		return nil
	}

	for _, pokemon := range conf.pokemonRegistry {
		fmt.Println(" -", pokemon.Name)
	}
	fmt.Println()
	return nil
}
