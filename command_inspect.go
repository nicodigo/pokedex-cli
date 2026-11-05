package main

import "fmt"

func commandInspect(conf *config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("must provide one (and only one) pokemon name")
	}

	pokemonName := args[0]

	pokemon, found := conf.pokemonRegistry[pokemonName]

	if !found {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokeType.Type.Name)
	}
	fmt.Println()
	return nil
}
