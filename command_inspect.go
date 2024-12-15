package main

import (
	"fmt"

	"github.com/abdullah0iq/pokedex/internal/pokeapi"
)

func commandInspect(conf *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("inspect command need 1 pokemon to be followed by the command\nUsage:inspect Pokemon-name")
	}
	if len(args) >= 2 {
		return fmt.Errorf("inspect command can't take more than one pokemon\nUsage:inspect Pokemon-name")
	}
	pokemonName := args[0]
	pokemon , ok := conf.pokeapiClient.Pokedex[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	inspectPokemon(pokemon)
	return nil
}	

func inspectPokemon(pokemon pokeapi.Pokemon) {
	p := fmt.Printf
	p("Name: %s\n", pokemon.Name)
	p("Height: %d\n", pokemon.Height)
	p("Weight: %d\n", pokemon.Weight)
	p("Stats:\n")
	for _, s := range pokemon.Stats {
		p("\t- %s: %v\n", s.Stat.Name, s.BaseStat)
	}
	p("Types:\n")
	for _, t := range pokemon.Types {
		p("\t- %s\n", t.Type.Name)
	}
}