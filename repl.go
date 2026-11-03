package main

import (
	"bufio"
	"fmt"
	"github.com/nicodigo/pokedex-cli/internal/pokeapi"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	commandRegistry map[string]cliCommand
	client          *pokeapi.PokeApiClient
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

		err := command.callback(conf)
		if err != nil {
			fmt.Printf("command: %s, error: %v\n", command.name, err)
		}

	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Print available commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Print a list of the next location areas",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Print a list of previous location areas",
			callback:    commandMapb,
		},
	}
}
