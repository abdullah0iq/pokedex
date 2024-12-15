package main

import "fmt"


func commandPokedex(conf *config, args []string) error {
	pokemons := conf.pokeapiClient.Pokedex
	if len(pokemons) == 0 {
		fmt.Println("You haven't got any pokemon, try catching them loser~")
		return nil
	}
	for _ , v := range pokemons {
		fmt.Printf(" -%s\n",v.Name)
	}
	return nil
}