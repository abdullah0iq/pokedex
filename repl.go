package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/abdullah0iq/pokedex/internal/pokeapi"
	"github.com/abdullah0iq/pokedex/internal/pokecache"
)

type config struct {
	pokeapiClient    pokeapi.Client
	pokeCache        pokecache.Cache
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			fmt.Println("0 input")
			continue
		}
		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			fmt.Println()
			continue
		}

	}
}

func cleanInput(text string) []string {

	trimmed := strings.TrimSpace(text)
	toLowerCase := strings.ToLower(trimmed)
	words := strings.Fields(toLowerCase)

	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
