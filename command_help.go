package main

import "fmt"

func commandHelp(cfg *config, arg []string) error {
	if len(arg) != 0 {
		fmt.Println("Exit command doesn't take any arguments")
		return nil
	}
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
