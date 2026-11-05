package main

import "fmt"

func commandHelp(conf *config, args []string) error {
	if len(conf.commandRegistry) == 0 {
		return fmt.Errorf("can't generate help, empty command registry")
	}
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range conf.commandRegistry{
		fmt.Printf("%s: %s\n", command.usage, command.description)
	}
	fmt.Println()
	return nil
}
