package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commandRegistry map[string]cliCommand

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	result := strings.Fields(lower)
	return result
}

func startRepl() {
	commandRegistry = map[string]cliCommand{
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
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		command, ok := commandRegistry[input[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback()
		if err != nil {
			fmt.Printf("command: %s, error: %v\n", command.name, err)
		}

	}
}
