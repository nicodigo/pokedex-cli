package main

import (
	"bufio"
	"fmt"
	"github.com/nicodigo/pokedex-cli/internal/pokeapi"
	"os"
	"strings"
)

type cliCommand struct {
	usage       string
	description string
	callback    func(*config, []string) error
}

type config struct {
	commandRegistry map[string]cliCommand
	client          *pokeapi.PokeApiClient
	pokemonRegistry map[string]pokeapi.Pokemon
	locationNext    *string
	locationPrev    *string
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	result := strings.Fields(lower)
	return result
}

func startRepl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		command, ok := conf.commandRegistry[input[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		args := input[1:]
		err := command.callback(conf, args)
		if err != nil {
			fmt.Printf("command: %s, error: %v\n", input[0], err)
		}

	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			usage:       "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			usage:       "help",
			description: "Prints available commands",
			callback:    commandHelp,
		},
		"map": {
			usage:       "map",
			description: "Prints a list of the next location areas",
			callback:    commandMapf,
		},
		"mapb": {
			usage:       "mapb",
			description: "Prints a list of previous location areas",
			callback:    commandMapb,
		},
		"explore": {
			usage:       "explore <location-name>",
			description: "Prints a list of all pokemons for a given area",
			callback:    commandExplore,
		},
		"catch": {
			usage:       "catch <pokemon-name>",
			description: "Tries to catch a pokemon and add it to the pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			usage:       "inspect <pokemon-name>",
			description: "Inspect a previously catched pokemon",
			callback:    commandInspect,
		},
	}
}
