package main

import (
	"time"

	"github.com/nicodigo/pokedex-cli/internal/pokeapi"
)

func main() {
	conf := config{
		commandRegistry: getCommands(),
		client:          pokeapi.NewPokeApiClient(time.Minute * 5),
		pokemonRegistry: map[string]pokeapi.Pokemon{},
		locationNext:            nil,
		locationPrev:            nil,
	}
	startRepl(&conf)
}
