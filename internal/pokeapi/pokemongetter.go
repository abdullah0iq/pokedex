package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) CatchPokemon(pokemonName string) (Pokemon , error) {
	pokemonUrl := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemonName)
	res , err := c.httpClient.Get(pokemonUrl)
	if err != nil {
		fmt.Printf("Couln't get '%s', check your spelling or this pokemon doesn't exist",pokemonName)
		return Pokemon{} , err
	}
	defer res.Body.Close()
	data , err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{} , err
	}
	pokemon := Pokemon{}
	if err = json.Unmarshal(data , &pokemon) ; err != nil{
		return Pokemon{} , err
	}
	return pokemon , nil
}