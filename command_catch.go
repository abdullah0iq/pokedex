package main

import (
	"fmt"
	"time"

	"github.com/abdullah0iq/pokedex/internal/pokeapi"
	"math/rand"
)

func commandCatch(conf *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Catch command need 1 location to be followed by the command\nUsage:Catch Pokemon-name")
	}
	if len(args) >= 2 {
		return fmt.Errorf("Catch command can't take more than one pokemon\nUsage:Catch Pokemon-name")
	}
	pokemonName := args[0]
	if _, ok := conf.pokeapiClient.Pokedex[pokemonName]; ok {
		fmt.Printf("You already have %s in your Pokedex!\n", pokemonName)
		return nil
	}
	pokemon, err := conf.pokeapiClient.CatchPokemon(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...", pokemonName)
	isCatched := catchThePokemon(pokemon)
	if isCatched {
		conf.pokeapiClient.Pokedex[pokemon.Name] = pokemon
		fmt.Printf("Congratulations! %s has been added to your Pokedex.\n", pokemon.Name)
	}
	return nil
}

func catchThePokemon(pokemon pokeapi.Pokemon) bool {
	baseExperience := pokemon.BaseExperience // Change this value to test with different Pokémon

	// MaxThreshold for determining difficulty (e.g., 300 is a high base experience)
	maxThreshold := 250.0

	// Calculate the success chance
	successChance := (maxThreshold - float64(baseExperience)) / maxThreshold

	// Seed the random number generator for different results each run
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 0 and 1
	randomValue := rand.Float64()


	// Determine if the Pokémon is caught
	if randomValue <= successChance {
		fmt.Println("You caught the Pokémon!")
		return true
	} else {
		fmt.Println("The Pokémon escaped!")
		return false
	}

}
