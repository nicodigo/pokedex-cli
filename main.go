package main

import "github.com/nicodigo/pokedex-cli/internal/pokeapi"

func main() {
	conf := config{
		commandRegistry: getCommands(),
		client:          pokeapi.NewPokeApiClient(),
		locationNext:            nil,
		locationPrev:            nil,
	}
	startRepl(&conf)
}
