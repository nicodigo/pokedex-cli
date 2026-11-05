package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(conf *config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("must provide one (and only one) pokemon name")
	}
	pokemon, err := conf.client.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	random := rand.Intn(pokemon.BaseExperience)

	if random <= 40 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		conf.pokemonRegistry[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}
