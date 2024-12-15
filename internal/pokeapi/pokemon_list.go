package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

func (c *Client) ListPokemons(locationName string) (LocationArea, error) {
	if locationName == "" {
		return LocationArea{}, fmt.Errorf("No location was inserted")
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", locationName)

	// check if the location cached
	if val, ok := c.cache.GetLocation(locationName); ok {
		location := LocationArea{}
		err := json.Unmarshal(val, &location)
		if err != nil {
			return LocationArea{}, err
		}
		return location, nil
	}

	// get the information from the api
	time.Sleep(3 * time.Second)
	res, err := c.httpClient.Get(url)
	if err != nil {

		return LocationArea{}, fmt.Errorf("Couldn't get the location information. Check your spelling")
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}
	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	c.cache.AddLocation(locationName, data)
	return locationArea, nil
}
