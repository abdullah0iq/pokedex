package main

import "fmt"

func commandExplore(conf *config , args []string) error{
	if len(args) == 0 {
		return fmt.Errorf("Explore command need 1 location to be follower by the command\nexplore location-name")
	}
	if len(args) >=2 {
		return fmt.Errorf("Explore command can't take more than one location\nexplore location-name")
	}

	location := args[0]
	fmt.Printf("Exploring %s...\n", location)
	locationArea ,err:= conf.pokeapiClient.ListPokemons(location)
	if err != nil {
		return err
	}
	for _ ,pokemon :=range locationArea.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}