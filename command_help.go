package main

import "fmt"

func commandHelp() error {
	if len(commandRegistry) == 0 {
		return fmt.Errorf("can't generate help, empty command registry")
	}
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range commandRegistry{
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}
